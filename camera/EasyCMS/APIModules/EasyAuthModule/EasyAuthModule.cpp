#include <stdio.h>

#include "EasyAuthModule.h"
#include "OSHeaders.h"
#include "QTSSModuleUtils.h"

// babosa add
#include "Task.h"
#include "OS.h"
#include "QTSServerInterface.h"
//add
#define __PLACEMENT_NEW_INLINE
#include"OSMapEx.h"
static OSMapEx sSessionIdMap;
//add

// STATIC VARIABLES
static QTSS_ModulePrefsObject sPrefs = NULL;
static QTSS_PrefsObject     sServerPrefs = NULL;
static QTSS_ServerObject    sServer = NULL;

// FUNCTION PROTOTYPES
static QTSS_Error   EasyAuthModuleDispatch(QTSS_Role inRole, QTSS_RoleParamPtr inParamBlock);
static QTSS_Error   Register(QTSS_Register_Params* inParams);
static QTSS_Error   Initialize(QTSS_Initialize_Params* inParams);
static QTSS_Error   RereadPrefs();
static QTSS_Error   MakeNonce(QTSS_Nonce_Params* inParams);//���������
static QTSS_Error   MakeAuth(QTSS_Nonce_Params* inParams);//��֤�����

//����Ƿ��г�ʱSessionID������
class SessionIDCheckTask : public Task
{
public:
	SessionIDCheckTask() : Task() { this->SetTaskName("SessionIDCheckTask"); this->Signal(Task::kStartEvent); }
	virtual ~SessionIDCheckTask() {}

private:
	virtual SInt64 Run();
};
static SessionIDCheckTask *pSessionIdTask = NULL;
//add 
SInt64 SessionIDCheckTask::Run()
{
	sSessionIdMap.CheckTimeoutAndDelete();//��鳬ʱ��SessionID������ɾ��
	return 60 * 1000;//һ����һ���
}
QTSS_Error EasyAuthModule_Main(void* inPrivateArgs)
{
	return _stublibrary_main(inPrivateArgs, EasyAuthModuleDispatch);
}

QTSS_Error  EasyAuthModuleDispatch(QTSS_Role inRole, QTSS_RoleParamPtr inParamBlock)
{
	switch (inRole)
	{
	case QTSS_Register_Role:
		return Register(&inParamBlock->regParams);
	case QTSS_Initialize_Role:
		return Initialize(&inParamBlock->initParams);
	case QTSS_RereadPrefs_Role:
		return RereadPrefs();
	case Easy_Nonce_Role:
		return MakeNonce(&inParamBlock->NonceParams);
	case Easy_Auth_Role:
		return MakeAuth(&inParamBlock->NonceParams);
	}
	return QTSS_NoErr;
}

QTSS_Error Register(QTSS_Register_Params* inParams)
{
	// Do role setup
	(void)QTSS_AddRole(QTSS_Initialize_Role);
	(void)QTSS_AddRole(QTSS_RereadPrefs_Role);
	(void)QTSS_AddRole(Easy_Nonce_Role);
	(void)QTSS_AddRole(Easy_Auth_Role);

	// Tell the server our name!
	static char* sModuleName = "EasyAuthModule";
	::strcpy(inParams->outModuleName, sModuleName);

	return QTSS_NoErr;
}

QTSS_Error Initialize(QTSS_Initialize_Params* inParams)
{
	QTSSModuleUtils::Initialize(inParams->inMessages, inParams->inServer, inParams->inErrorLogStream);
	sServer = inParams->inServer;
	sServerPrefs = inParams->inPrefs;
	sPrefs = QTSSModuleUtils::GetModulePrefsObject(inParams->inModule);


	//string strTest("123456");
	//sSessionIdMap.Insert(strTest);
	pSessionIdTask = new SessionIDCheckTask();//add,���SessionID�Ƿ�ʱ��TASK
	return RereadPrefs();
}

QTSS_Error RereadPrefs()
{
	return QTSS_NoErr;
}
QTSS_Error MakeNonce(QTSS_Nonce_Params* inParams)//���������
{
	string strTemp = sSessionIdMap.GererateAndInsert();
	strcpy(inParams->pNonce, strTemp.c_str());
	return QTSS_NoErr;
}
QTSS_Error MakeAuth(QTSS_Nonce_Params* inParams)//���������
{
	string strSessionID(inParams->pNonce);
	*(inParams->pResult) = (char)sSessionIdMap.FindAndDelete(strSessionID);
	return QTSS_NoErr;
}