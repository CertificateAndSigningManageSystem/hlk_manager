/*
 * Copyright (c) 2023 ivfzhou
 * hlk_manager is licensed under Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at:
 *          http://license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
 * EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
 * MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 */

package service

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"time"

	"gitee.com/CertificateAndSigningManageSystem/common/ctxs"
	"gitee.com/CertificateAndSigningManageSystem/common/log"
	"gitee.com/CertificateAndSigningManageSystem/common/model"
	"gitee.com/CertificateAndSigningManageSystem/hlk_manager/consts"
)

// 测试任务信息
type testJobInfo struct {
	Id             int    `json:"Id,omitempty"`
	FileName       string `json:"FileName,omitempty"`
	Driver         string `json:"Driver,omitempty"`
	UserName       string `json:"UserName,omitempty"`
	ServiceName    string `json:"ServiceName,omitempty"`
	DriverFileID   string `json:"DriverFileID,omitempty"`
	DriverFilePath string `json:"DriverFilePath,omitempty"`
	whqlConfig     `json:"whqlConfig,omitempty"`
}

// 测试配置
type whqlConfig struct {
	elamConfig                            `json:"elamConfig,omitempty"`
	wfpConfig                             `json:"wfpConfig,omitempty"`
	IsWindowsDriverProject                bool `json:"IsWindowsDriverProject,omitempty"`
	IsHSPCompatibility                    bool `json:"IsHSPCompatibility"`
	AudioCodecVerifyAudioEffectsDiscovery bool `json:"AudioCodecVerifyAudioEffectsDiscovery,omitempty"`
}

type elamConfig struct {
	IsWdBootMVIMember int `json:"isWdBootMVIMember,omitempty"`
}

type wfpConfig struct {
	ProductName                       string `json:"productName,omitempty"`
	EnableDriverVerifier              int    `json:"enableDriverVerifier,omitempty"`
	CalloutDriver                     int    `json:"calloutDriver,omitempty"`
	IsAFirewall                       int    `json:"isAFirewall,omitempty"`
	LayeredOnMicrosoftWindowsFirewall int    `json:"layeredOnMicrosoftWindowsFirewall,omitempty"`
	DoesMACFiltering                  int    `json:"doesMACFiltering,omitempty"`
	DoesVSwitchFiltering              int    `json:"doesVSwitchFiltering,omitempty"`
	DoesPacketInjection               int    `json:"doesPacketInjection,omitempty"`
	DoesStreamInjection               int    `json:"doesStreamInjection,omitempty"`
	DoesConnectionProxying            int    `json:"doesConnectionProxying,omitempty"`
	SupportModernApplications         int    `json:"supportModernApplications,omitempty"`
	CleanUninstall                    int    `json:"cleanUninstall,omitempty"`
	NoProxyDeadlocks                  int    `json:"noProxyDeadlocks,omitempty"`
	IdentifyingProvider               int    `json:"identifyingProvider,omitempty"`
	AssociateProvider                 int    `json:"associateProvider,omitempty"`
	TerminatingFilter                 int    `json:"terminatingFilter,omitempty"`
	UseOwnSubLayer                    int    `json:"useOwnSubLayer,omitempty"`
	MaintainHelperClass               int    `json:"maintainHelperClass,omitempty"`
	NoAVs                             int    `json:"noAVs,omitempty"`
	NonTampered3rdPartyObjects        int    `json:"nonTampered3rdPartyObjects,omitempty"`
	NoPacketInjectionDeadlocks        int    `json:"noPacketInjectionDeadlocks,omitempty"`
	NoStreamStarvation                int    `json:"noStreamStarvation,omitempty"`
	SupportPowerManagement            int    `json:"supportPowerManagement,omitempty"`
	WFPObjectEnumAndACLs              int    `json:"wfpObjectEnumAndACLs,omitempty"`
	MaxWinSock                        int    `json:"maxWinSock,omitempty"`
	ProperlyDisableWindowsFirewall    int    `json:"properlyDisableWindowsFirewall,omitempty"`
	NoPermitBlockAll                  int    `json:"noPermitBlockAll,omitempty"`
	SupportTupleExceptions            int    `json:"supportTupleExceptions,omitempty"`
	SupportAppExceptions              int    `json:"supportAppExceptions,omitempty"`
	SupportMACAddressExceptions       int    `json:"supportMACAddressExceptions,omitempty"`
	UseWFP                            int    `json:"useWFP,omitempty"`
	SupportARP                        int    `json:"supportARP,omitempty"`
	SupportNeighborDiscovery          int    `json:"supportNeighborDiscovery,omitempty"`
	SupportDHCP                       int    `json:"supportDHCP,omitempty"`
	SupportIPv4                       int    `json:"supportIPv4,omitempty"`
	SupportIPv6                       int    `json:"supportIPv6,omitempty"`
	SupportDNS                        int    `json:"supportDNS,omitempty"`
	Support6To4                       int    `json:"support6To4,omitempty"`
	SupportAutomaticUpdates           int    `json:"supportAutomaticUpdates,omitempty"`
	SupportBasicWebsiteBrowsing       int    `json:"supportBasicWebsiteBrowsing,omitempty"`
	SupportFileAndPrinterSharing      int    `json:"supportFileAndPrinterSharing,omitempty"`
	SupportICMPErrorMesages           int    `json:"supportICMPErrorMesages,omitempty"`
	SupportInternetStreaming          int    `json:"supportInternetStreaming,omitempty"`
	SupportMediaExtenderStreaming     int    `json:"supportMediaExtenderStreaming,omitempty"`
	SupportMobileBroadBand            int    `json:"supportMobileBroadBand,omitempty"`
	SupportPeerNameResolution         int    `json:"supportPeerNameResolution,omitempty"`
	SupportRemoteAssistance           int    `json:"supportRemoteAssistance,omitempty"`
	SupportRemoteDesktop              int    `json:"supportRemoteDesktop,omitempty"`
	SupportTeredo                     int    `json:"supportTeredo,omitempty"`
	SupportVirtualPrivateNetworking   int    `json:"supportVirtualPrivateNetworking,omitempty"`
	InteropWithOtherExtensions        int    `json:"interopWithOtherExtensions,omitempty"`
	NoEgressModification              int    `json:"noEgressModification,omitempty"`
	SupportLiveMigration              int    `json:"supportLiveMigration,omitempty"`
	SupportRemoval                    int    `json:"supportRemoval,omitempty"`
	SupportReordering                 int    `json:"supportReordering,omitempty"`
	InteropWithWFPSampler             int    `json:"interopWithWFPSampler,omitempty"`
}

// StartTestSystemHandler 运行测试机
func StartTestSystemHandler(ctx context.Context) {

}

func handleTestSystemJob(ctx context.Context) {
	defer func() {
		if p := recover(); p != nil {
			log.Error(ctx, p)
		}
		handleTestSystemJob(ctx)
	}()

	for {
		ctx := ctxs.NewCtx("handleTestSystemJob")
		time.Sleep(10 * time.Second)

		// 获取可能正在测试中的任务信息
		job, err := readTestingJob(ctx)
		if err != nil {
			log.Error(ctx, err)
			continue
		}
		// 检查任务状态
		if job.Id > 0 {
			jobInfo, err := queryJobInfo(ctx, job.Id)
			if err != nil {
				log.Error(ctx, err)
				continue
			}
			// 任务非法，回滚测试机器
			if jobInfo.Id <= 0 || jobInfo.Status != model.TWinSignJob_Status_NeedHLKTest {

			}
		}
	}
}

func notifyRestoreVM(ctx context.Context) {

}

// 读取硬盘，获取可能正在测试中的任务信息
func readTestingJob(ctx context.Context) (*testJobInfo, error) {
	// 读取文件
	file, err := os.ReadFile(consts.JobInfoPath)
	if err != nil {
		// 没有在测试中的任务
		if errors.Is(err, os.ErrNotExist) {
			return &testJobInfo{}, nil
		}
		log.Error(ctx, err)
		return nil, err
	}

	// 反序列化
	var info testJobInfo
	if err = json.Unmarshal(file, &info); err != nil {
		log.Error(ctx, err)
		return nil, err
	}

	return &info, nil
}
