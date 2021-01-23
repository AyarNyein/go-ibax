// +build windows

/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package system

import (
	"os"
)

/*
#include <windows.h>
#include <stdio.h>
#include <stdlib.h>
#include <TlHelp32.h>

void kill_childproc( DWORD myprocID) {
	PROCESSENTRY32 pe;

	memset(&pe, 0, sizeof(PROCESSENTRY32));
	pe.dwSize = sizeof(PROCESSENTRY32);

	HANDLE hSnap = CreateToolhelp32Snapshot(TH32CS_SNAPPROCESS, 0);
	if (Process32First(hSnap, &pe))
	{
	    BOOL bContinue = TRUE;

	    while (bContinue)
	        }
	        bContinue = Process32Next(hSnap, &pe);
	    }
	}
}
*/
import "C"

// lstrcmp( pe.szExeFile, TEXT("tmp_daylight.exe")) != 0 && lstrcmp( pe.szExeFile, TEXT("daylight.exe")) != 0

func killChildProc() {
	C.kill_childproc(C.DWORD(os.Getpid()))
}
