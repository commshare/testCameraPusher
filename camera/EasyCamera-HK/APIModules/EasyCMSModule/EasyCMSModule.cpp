/*
	Copyright (c) 2012-2016 EasyDarwin.ORG.  All rights reserved.
	Github: https://github.com/EasyDarwin
	WEChat: EasyDarwin
	Website: http://www.easydarwin.org
*/
/*
	File:       EasyCMSModule.cpp
	Contains:   Implementation of EasyCMSModule class.
*/

#include "EasyCMSModule.h"
#include "QTSSModuleUtils.h"
#include "QTSServerInterface.h"
#include "EasyCMSSession.h"

// STATIC DATA
static QTSS_PrefsObject				sServerPrefs = NULL;	//������������
static QTSS_ServerObject			sServer = NULL;	//����������
static QTSS_ModulePrefsObject		sEasyCMSModulePrefs = NULL;	//��ǰģ������

static EasyCMSSession*				sCMSSession = NULL; //ΨһEasyCMSSession����

// FUNCTION PROTOTYPES
static QTSS_Error EasyCMSModuleDispatch(QTSS_Role inRole, QTSS_RoleParamPtr inParams);
static QTSS_Error Register_EasyCMSModule(QTSS_Register_Params* inParams);
static QTSS_Error Initialize_EasyCMSModule(QTSS_Initialize_Params* inParams);
static QTSS_Error RereadPrefs_EasyCMSModule();

// FUNCTION IMPLEMENTATIONS
QTSS_Error EasyCMSModule_Main(void* inPrivateArgs)
{
	return _stublibrary_main(inPrivateArgs, EasyCMSModuleDispatch);
}

QTSS_Error EasyCMSModuleDispatch(QTSS_Role inRole, QTSS_RoleParamPtr inParams)
{
	switch (inRole)
	{
	case QTSS_Register_Role:
		return Register_EasyCMSModule(&inParams->regParams);
	case QTSS_Initialize_Role:
		return Initialize_EasyCMSModule(&inParams->initParams);
	case QTSS_RereadPrefs_Role:
		return RereadPrefs_EasyCMSModule();
	}
	return QTSS_NoErr;
}

QTSS_Error Register_EasyCMSModule(QTSS_Register_Params* inParams)
{
	// Do role & attribute setup
	(void)QTSS_AddRole(QTSS_Initialize_Role);
	(void)QTSS_AddRole(QTSS_RereadPrefs_Role);

	// Tell the server our name!
	static char* sModuleName = "EasyCMSModule";
	::strcpy(inParams->outModuleName, sModuleName);

	return QTSS_NoErr;
}

QTSS_Error Initialize_EasyCMSModule(QTSS_Initialize_Params* inParams)
{
	// Setup module utils
	QTSSModuleUtils::Initialize(inParams->inMessages, inParams->inServer, inParams->inErrorLogStream);

	// Setup global data structures
	sServerPrefs = inParams->inPrefs;
	sServer = inParams->inServer;
	sEasyCMSModulePrefs = QTSSModuleUtils::GetModulePrefsObject(inParams->inModule);

	//��ȡEasyCMSModule����
	RereadPrefs_EasyCMSModule();

	EasyCMSSession::Initialize(sEasyCMSModulePrefs);

	//��������ʼEasyCMSSession����
	sCMSSession = new EasyCMSSession();
	sCMSSession->Signal(Task::kStartEvent);

	return QTSS_NoErr;
}

QTSS_Error RereadPrefs_EasyCMSModule()
{
	return QTSS_NoErr;
}