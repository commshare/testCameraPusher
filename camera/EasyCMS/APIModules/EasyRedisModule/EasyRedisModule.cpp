#include <stdio.h>
#include "EasyRedisModule.h"
#include "OSHeaders.h"
#include "QTSSModuleUtils.h"
#include "EasyRedisClient.h"
#include "QTSServerInterface.h"
#include "HTTPSessionInterface.h"
#include "Format.h"

// STATIC VARIABLES
static QTSS_ModulePrefsObject	modulePrefs		= NULL;
static QTSS_PrefsObject			sServerPrefs    = NULL;
static QTSS_ServerObject		sServer			= NULL;

// Redis IP
static char*            sRedis_IP				= NULL;
static char*            sDefaultRedis_IP_Addr	= "127.0.0.1";
// Redis Port
static UInt16			sRedisPort				= 6379;
static UInt16			sDefaultRedisPort		= 6379;
// Redis user
static char*            sRedisUser				= NULL;
static char*            sDefaultRedisUser		= "admin";
// Redis password
static char*            sRedisPassword			= NULL;
static char*            sDefaultRedisPassword	= "admin";
// EasyCMS
static char*			sCMSIP					= NULL;
static UInt16			sCMSPort				= 10000;
static EasyRedisClient* sRedisClient			= NULL;//the object pointer that package the redis operation
static bool				sIfConSucess			= false;
static OSMutex			sMutex;

// FUNCTION PROTOTYPES
static QTSS_Error   EasyRedisModuleDispatch(QTSS_Role inRole, QTSS_RoleParamPtr inParamBlock);
static QTSS_Error   Register(QTSS_Register_Params* inParams);
static QTSS_Error   Initialize(QTSS_Initialize_Params* inParams);
static QTSS_Error   RereadPrefs();

static QTSS_Error RedisConnect();
static QTSS_Error RedisInit();
static QTSS_Error RedisTTL();
static QTSS_Error RedisAddDevName(QTSS_StreamName_Params* inParams);
static QTSS_Error RedisDelDevName(QTSS_StreamName_Params* inParams);
static QTSS_Error RedisGetAssociatedDarwin(QTSS_GetAssociatedDarwin_Params* inParams);
static QTSS_Error RedisGetBestDarwin(QTSS_GetBestDarwin_Params * inParams);
static QTSS_Error RedisGenStreamID(QTSS_GenStreamID_Params* inParams);

QTSS_Error EasyRedisModule_Main(void* inPrivateArgs)
{
	return _stublibrary_main(inPrivateArgs, EasyRedisModuleDispatch);
}

QTSS_Error  EasyRedisModuleDispatch(QTSS_Role inRole, QTSS_RoleParamPtr inParamBlock)
{
	switch (inRole)
	{
	case QTSS_Register_Role:
		return Register(&inParamBlock->regParams);
	case QTSS_Initialize_Role:
		return Initialize(&inParamBlock->initParams);
	case QTSS_RereadPrefs_Role:
		return RereadPrefs();
	case Easy_RedisAddDevName_Role:
		return RedisAddDevName(&inParamBlock->StreamNameParams);
	case Easy_RedisDelDevName_Role:
		return RedisDelDevName(&inParamBlock->StreamNameParams);
	case Easy_RedisTTL_Role:
		return RedisTTL();
	case Easy_RedisGetEasyDarwin_Role:
		return RedisGetAssociatedDarwin(&inParamBlock->GetAssociatedDarwinParams);
	case Easy_RedisGetBestEasyDarwin_Role:
		return RedisGetBestDarwin(&inParamBlock->GetBestDarwinParams);
	case Easy_RedisGenStreamID_Role:
		return RedisGenStreamID(&inParamBlock->GenStreamIDParams);
	}
	return QTSS_NoErr;
}

QTSS_Error Register(QTSS_Register_Params* inParams)
{
	// Do role setup
	(void)QTSS_AddRole(QTSS_Initialize_Role);
	(void)QTSS_AddRole(QTSS_RereadPrefs_Role);
	(void)QTSS_AddRole(Easy_RedisTTL_Role);
	(void)QTSS_AddRole(Easy_RedisAddDevName_Role);
	(void)QTSS_AddRole(Easy_RedisDelDevName_Role);
	(void)QTSS_AddRole(Easy_RedisGetEasyDarwin_Role);
	(void)QTSS_AddRole(Easy_RedisGetBestEasyDarwin_Role);
	(void)QTSS_AddRole(Easy_RedisGenStreamID_Role);
	// Tell the server our name!
	static char* sModuleName = "EasyRedisModule";
	::strcpy(inParams->outModuleName, sModuleName);

	return QTSS_NoErr;
}

QTSS_Error Initialize(QTSS_Initialize_Params* inParams)
{
	QTSSModuleUtils::Initialize(inParams->inMessages, inParams->inServer, inParams->inErrorLogStream);
	sServer = inParams->inServer;
	sServerPrefs = inParams->inPrefs;
	modulePrefs = QTSSModuleUtils::GetModulePrefsObject(inParams->inModule);

	RereadPrefs();

	sRedisClient = new EasyRedisClient();

	RedisConnect();

	return QTSS_NoErr;
}

QTSS_Error RereadPrefs()
{
	delete[] sRedis_IP;
	sRedis_IP = QTSSModuleUtils::GetStringAttribute(modulePrefs, "redis_ip", sDefaultRedis_IP_Addr);

	QTSSModuleUtils::GetAttribute(modulePrefs, "redis_port", qtssAttrDataTypeUInt16, &sRedisPort, &sDefaultRedisPort, sizeof(sRedisPort));

	delete[] sRedisUser;
	sRedisUser = QTSSModuleUtils::GetStringAttribute(modulePrefs, "redis_user", sDefaultRedisUser);

	delete[] sRedisPassword;
	sRedisPassword = QTSSModuleUtils::GetStringAttribute(modulePrefs, "redis_password", sDefaultRedisPassword);

	//get cms ip and port
	delete[] sCMSIP;
	(void)QTSS_GetValueAsString(sServerPrefs, qtssPrefsMonitorWANIPAddr, 0, &sCMSIP);

	UInt32 len = sizeof(SInt32);
	(void)QTSS_GetValue(sServerPrefs, qtssPrefsMonitorWANPort, 0, (void*)&sCMSPort, &len);

	return QTSS_NoErr;
}

QTSS_Error RedisConnect()
{
	if (sIfConSucess)
		return QTSS_NoErr;

	std::size_t timeout = 1;//timeout second
	if (sRedisClient->ConnectWithTimeOut(sRedis_IP, sRedisPort, timeout) == EASY_REDIS_OK)//return 0 if connect sucess
	{
		qtss_printf("Connect redis sucess\n");
		sIfConSucess = true;
		std::size_t timeoutSocket = 1;//timeout socket second
		sRedisClient->SetTimeout(timeoutSocket);
		RedisInit();
	}
	else
	{
		qtss_printf("Connect redis failed\n");
		sIfConSucess = false;
	}
	return (QTSS_Error)(!sIfConSucess);
}

QTSS_Error RedisInit()//only called by RedisConnect after connect redis sucess
{
	//ÿһ����redis���Ӻ󣬶�Ӧ�������һ�ε����ݴ洢��ʹ�ø��ǻ���ֱ������ķ�ʽ,��������ʹ�ù��߸��Ӹ�Ч
	char chTemp[128] = { 0 };

	do
	{
		//1,redis������֤
		sprintf(chTemp, "auth %s", sRedisPassword);
		sRedisClient->AppendCommand(chTemp);

		//2,CMSΨһ��Ϣ�洢(������һ�εĴ洢)
		sprintf(chTemp, "sadd EasyCMSName %s:%d", sCMSIP, sCMSPort);
		sRedisClient->AppendCommand(chTemp);


		//3,CMS���Դ洢,���ö��filedʹ��hmset������ʹ��hset(������һ�εĴ洢)
		sprintf(chTemp, "hmset %s:%d_Info IP %s PORT %d", sCMSIP, sCMSPort, sCMSIP, sCMSPort);
		sRedisClient->AppendCommand(chTemp);

		//4,����豸���ƴ洢����Ϊ����֮ǰ������֮����豸����һ��÷����˱仯����˱�����ִ���������
		sprintf(chTemp, "del %s:%d_DevName", sCMSIP, sCMSPort);
		sRedisClient->AppendCommand(chTemp);

		OSRefTableEx*  deviceRefTable = QTSServerInterface::GetServer()->GetDeviceSessionMap();
		OSMutex *mutexMap = deviceRefTable->GetMutex();
		OSHashMap  *deviceMap = deviceRefTable->GetMap();
		OSRefIt itRef;
		string strAllDevices;
		{
			OSMutexLocker lock(mutexMap);
			for (itRef = deviceMap->begin(); itRef != deviceMap->end(); itRef++)
			{
				strDevice *deviceInfo = (((HTTPSessionInterface*)(itRef->second->GetObjectPtr()))->GetDeviceInfo());
				strAllDevices = strAllDevices + ' ' + deviceInfo->serial_;
			}
		}

		char *chNewTemp = new char[strAllDevices.size() + 128];//ע�⣬���ﲻ����ʹ��chTemp����Ϊ���Ȳ�ȷ�������ܵ��»��������
		//5,�豸���ƴ洢
		sprintf(chNewTemp, "sadd %s:%d_DevName%s", sCMSIP, sCMSPort, strAllDevices.c_str());
		sRedisClient->AppendCommand(chNewTemp);
		delete[] chNewTemp;

		//6,�������15�룬��֮��ǰCMS�Ѿ���ʼ�ṩ������
		sprintf(chTemp, "setex %s:%d_Live 15 1", sCMSIP, sCMSPort);
		sRedisClient->AppendCommand(chTemp);

		bool bBreak = false;
		easyRedisReply* reply = NULL;
		for (int i = 0; i < 6; i++)
		{
			if (EASY_REDIS_OK != sRedisClient->GetReply((void**)&reply))
			{
				bBreak = true;
				if (reply)
					EasyFreeReplyObject(reply);
				break;
			}
			EasyFreeReplyObject(reply);
		}
		if (bBreak)//˵��redisGetReply�����˴���
			break;
		return QTSS_NoErr;
	} while (0);
	//�ߵ���˵�������˴�����Ҫ��������,������������һ��ִ������ʱ����,����������ñ�־λ
	sRedisClient->Free();

	sIfConSucess = false;
	return (QTSS_Error)false;
}

QTSS_Error RedisAddDevName(QTSS_StreamName_Params* inParams)
{
	OSMutexLocker mutexLock(&sMutex);
	if (!sIfConSucess)
		return QTSS_NotConnected;

	char chKey[128] = { 0 };
	sprintf(chKey, "%s:%d_DevName", sCMSIP, sCMSPort);

	int ret = sRedisClient->SAdd(chKey, inParams->inStreamName);
	if (ret == -1)//fatal err,need reconnect
	{
		sRedisClient->Free();
		sIfConSucess = false;
	}

	return ret;
}

QTSS_Error RedisDelDevName(QTSS_StreamName_Params* inParams)
{
	OSMutexLocker mutexLock(&sMutex);
	if (!sIfConSucess)
		return QTSS_NotConnected;

	char chKey[128] = { 0 };
	sprintf(chKey, "%s:%d_DevName", sCMSIP, sCMSPort);

	int ret = sRedisClient->SRem(chKey, inParams->inStreamName);
	if (ret == -1)//fatal err,need reconnect
	{
		sRedisClient->Free();
		sIfConSucess = false;
	}

	return ret;
}

QTSS_Error RedisTTL()//ע�⵱������һ��ʱ��ܲ�ʱ���ܻ���Ϊ��ʱʱ��ﵽ������key��ɾ������ʱӦ���������ø�key
{

	OSMutexLocker mutexLock(&sMutex);

	if (RedisConnect() != QTSS_NoErr)//ÿһ��ִ������֮ǰ��������redis,�����ǰredis��û�гɹ�����
		return QTSS_NotConnected;

	char chKey[128] = { 0 };//ע��128λ�Ƿ��㹻
	sprintf(chKey, "%s:%d_Live 15", sCMSIP, sCMSPort);//���ĳ�ʱʱ��

	int ret = sRedisClient->SetExpire(chKey, 15);
	if (ret == -1)//fatal error
	{
		sRedisClient->Free();
		sIfConSucess = false;
		return QTSS_NotConnected;
	}
	else if (ret == 1)
	{
		return QTSS_NoErr;
	}
	else if (ret == 0)//the key doesn't exist, reset
	{
		sprintf(chKey, "%s:%d_Live", sCMSIP, sCMSPort);
		int retret = sRedisClient->SetEX(chKey, 15, "1");
		if (retret == -1)//fatal error
		{
			sRedisClient->Free();
			sIfConSucess = false;
		}
		return retret;
	}
	else
	{
		return ret;
	}
}

QTSS_Error RedisGetAssociatedDarwin(QTSS_GetAssociatedDarwin_Params* inParams)
{
	OSMutexLocker mutexLock(&sMutex);

	if (!sIfConSucess)
		return QTSS_NotConnected;

	string strPushName = Format("%s/%s", string(inParams->inSerial), string(inParams->inChannel));

	//1. get the list of EasyDarwin
	easyRedisReply * reply = (easyRedisReply *)sRedisClient->SMembers("EasyDarwinName");
	if (reply == NULL)
	{
		sRedisClient->Free();
		sIfConSucess = false;
		return QTSS_NotConnected;
	}

	//2.judge if the EasyDarwin is ilve and contain serial/channel.sdp
	if ((reply->elements > 0) && (reply->type == EASY_REDIS_REPLY_ARRAY))
	{
		easyRedisReply* childReply = NULL;
		for (size_t i = 0; i < reply->elements; i++)
		{
			childReply = reply->element[i];
			string strChileReply(childReply->str);

			string strTemp = Format("exists %s", strChileReply + "_Live");
			sRedisClient->AppendCommand(strTemp.c_str());

			strTemp = Format("sismember %s %s", strChileReply + "_PushName", strPushName);
			sRedisClient->AppendCommand(strTemp.c_str());
		}

		easyRedisReply *reply2 = NULL, *reply3 = NULL;
		for (size_t i = 0; i < reply->elements; i++)
		{
			if (sRedisClient->GetReply((void**)&reply2) != EASY_REDIS_OK)
			{
				EasyFreeReplyObject(reply);
				if (reply2)
				{
					EasyFreeReplyObject(reply2);
				}
				sRedisClient->Free();
				sIfConSucess = false;
				return QTSS_NotConnected;
			}
			if (sRedisClient->GetReply((void**)&reply3) != EASY_REDIS_OK)
			{
				EasyFreeReplyObject(reply);
				if (reply3)
				{
					EasyFreeReplyObject(reply3);
				}
				sRedisClient->Free();
				sIfConSucess = false;
				return QTSS_NotConnected;
			}

			if ((reply2->type == EASY_REDIS_REPLY_INTEGER) && (reply2->integer == 1) &&
				(reply3->type == EASY_REDIS_REPLY_INTEGER) && (reply3->integer == 1))
			{//find it
				string strIpPort(reply->element[i]->str);
				int ipos = strIpPort.find(':');//judge error
				memcpy(inParams->outDssIP, strIpPort.c_str(), ipos);
				memcpy(inParams->outDssPort, &strIpPort[ipos + 1], strIpPort.size() - ipos - 1);
				//break;//can't break,as 1 to 1
			}
			EasyFreeReplyObject(reply2);
			EasyFreeReplyObject(reply3);
		}
	}
	EasyFreeReplyObject(reply);
	return QTSS_NoErr;
}

QTSS_Error RedisGetBestDarwin(QTSS_GetBestDarwin_Params * inParams)
{
	OSMutexLocker mutexLock(&sMutex);

	QTSS_Error theErr = QTSS_NoErr;

	if (!sIfConSucess)
		return QTSS_NotConnected;


	char chTemp[128] = { 0 };

	//1. get the list of EasyDarwin
	easyRedisReply * reply = (easyRedisReply *)sRedisClient->SMembers("EasyDarwinName");
	if (reply == NULL)
	{
		sRedisClient->Free();
		sIfConSucess = false;
		return QTSS_NotConnected;
	}

	//2.judge if the EasyDarwin is ilve and get the RTP
	if ((reply->elements > 0) && (reply->type == EASY_REDIS_REPLY_ARRAY))
	{
		easyRedisReply* childReply = NULL;
		for (size_t i = 0; i < reply->elements; i++)
		{
			childReply = reply->element[i];
			string strChileReply(childReply->str);

			sprintf(chTemp, "exists %s", (strChileReply + "_Live").c_str());
			sRedisClient->AppendCommand(chTemp);

			sprintf(chTemp, "hget %s %s", (strChileReply + "_Info").c_str(), "RTP");
			sRedisClient->AppendCommand(chTemp);
		}

		int key = -1, keynum = 0;
		easyRedisReply *reply2 = NULL, *reply3 = NULL;
		for (size_t i = 0; i < reply->elements; i++)
		{
			if (sRedisClient->GetReply((void**)&reply2) != EASY_REDIS_OK)
			{
				EasyFreeReplyObject(reply);
				if (reply2)
				{
					EasyFreeReplyObject(reply2);
				}
				sRedisClient->Free();
				sIfConSucess = false;
				return QTSS_NotConnected;
			}

			if (sRedisClient->GetReply((void**)&reply3) != EASY_REDIS_OK)
			{
				EasyFreeReplyObject(reply);
				if (reply3)
				{
					EasyFreeReplyObject(reply3);
				}
				sRedisClient->Free();
				sIfConSucess = false;
				return QTSS_NotConnected;
			}

			if ((reply2->type == EASY_REDIS_REPLY_INTEGER) && (reply2->integer == 1) &&
				(reply3->type == EASY_REDIS_REPLY_STRING))
			{//find it
				int RTPNum = atoi(reply3->str);
				if (key == -1)
				{
					key = i;
					keynum = RTPNum;
				}
				else
				{
					if (RTPNum < keynum)//find better
					{
						key = i;
						keynum = RTPNum;
					}
				}
			}
			EasyFreeReplyObject(reply2);
			EasyFreeReplyObject(reply3);
		}
		if (key == -1)//no one live
		{
			theErr = QTSS_Unimplemented;
		}
		else
		{
			string strIpPort(reply->element[key]->str);
			int ipos = strIpPort.find(':');//judge error
			memcpy(inParams->outDssIP, strIpPort.c_str(), ipos);
			memcpy(inParams->outDssPort, &strIpPort[ipos + 1], strIpPort.size() - ipos - 1);
		}
	}
	else//û�п��õ�EasyDarWin
	{
		theErr = QTSS_Unimplemented;
	};
	EasyFreeReplyObject(reply);
	return theErr;
}

QTSS_Error RedisGenStreamID(QTSS_GenStreamID_Params* inParams)
{
	//�㷨���٣��������sessionID����redis���Ƿ��д洢��û�оʹ���redis�ϣ��еĻ��������ɣ�ֱ��û��Ϊֹ
	OSMutexLocker mutexLock(&sMutex);

	if (!sIfConSucess)
		return QTSS_NotConnected;

	easyRedisReply* reply = NULL;
	char chTemp[128] = { 0 };
	string strSessioionID;

	do
	{
		if (reply)//�ͷ���һ����Ӧ
			EasyFreeReplyObject(reply);

		strSessioionID = OSMapEx::GenerateSessionIdForRedis(sCMSIP, sCMSPort);

		sprintf(chTemp, "SessionID_%s", strSessioionID.c_str());
		reply = (easyRedisReply*)sRedisClient->Exists(chTemp);
		if (NULL == reply)//������Ҫ��������
		{
			sRedisClient->Free();
			sIfConSucess = false;
			return QTSS_NotConnected;
		}
	} while ((reply->type == EASY_REDIS_REPLY_INTEGER) && (reply->integer == 1));
	EasyFreeReplyObject(reply);//�ͷ����һ���Ļ�Ӧ

	//�ߵ���˵���ҵ���һ��Ψһ��SessionID�����ڽ����洢��redis��
	sprintf(chTemp, "SessionID_%s", strSessioionID.c_str());//�߼��汾֧��setpx�����ó�ʱʱ��Ϊms
	if (sRedisClient->SetEX(chTemp, inParams->inTimeoutMil / 1000, "1") == -1)
	{
		sRedisClient->Free();
		sIfConSucess = false;
		return QTSS_NotConnected;
	}
	strcpy(inParams->outStreanID, strSessioionID.c_str());
	return QTSS_NoErr;
}