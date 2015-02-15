// generated by stringer -type=ReplyCode; DO NOT EDIT

package irc

import "fmt"

const _ReplyCode_name = "ReplyWelcomeReplyYourHostReplyCreatedReplyMyInfoReplyBounceReplyTraceLinkReplyTraceConnectingReplyTraceHandshakeReplyTraceUnknownReplyTraceOperatorReplyTraceUserReplyTraceServerReplyTraceServiceReplyTraceNewTypeReplyTraceClassReplyTraceReconnectReplyStatsLinkInfoReplyStatsCommandsReplyEndOfStatsReplyUserModeIsReplyServerListReplyServerListEndReplyStatsUpTimeReplyStatsOLineReplyLUserClientReplyLUserOpReplyLUserUnknownReplyLUserChannelsReplyLUserMeReplyAdminMeReplyAdminLoc1ReplyAdminLoc2ReplyAdminEmailReplyTraceLogReplyTraceEndReplyTryAgainReplyAwayReplyUserHostReplyISONReplyUnAwayReplyNowAwayReplyWhoIsUserReplyWhoIsServerReplyWhoIsOperatorReplyWhoWasUserReplyEndOfWhoReplyWhoIsIdleReplyEndOfWhoIsReplyWhoIsChannelsReplyListStartReplyListReplyListEndReplyChannelModeIsReplyUniqueOpIsReplyNoTopicReplyTopicReplyInvitingReplySummoningReplyInviteListReplyEndOfInviteListReplyExceptListReplyEndOfExceptListReplyVersionReplyWhoReplyReplyNameReplyReplyLinksReplyEndOfLinksReplyEndOfNamesReplyBanListReplyEndOfBanListReplyEndOfWhoWasReplyInfoReplyMOTDReplyEndOfInfoReplyMOTDStartReplyEndOfMOTDReplyYoureOperReplyRehashingReplyYoureServiceReplyTimeReplyUserStartReplyUsersReplyEndOfUsersReplyNoUsersErrorNoSuchNicknameErrorNoSuchServerErrorNoSuchChannelErrorCannotSendToChannelErrorTooManyChannelsErrorWasNoSuchNicknameErrorTooManyTargetsErrorNoSuchServiceErrorNoOriginErrorNoRecipientErrorNoTextToSendErrorNoTopLevelErrorWildTopLevelErrorBadMaskErrorUnknownCommandErrorNoMOTDErrorNoAdminInfoErrorFileErrorErrorNoNicknameGivenErrorErroneousNicknameErrorNicknameInUseErrorNicknameCollisionErrorUnavailableResourceErrorUserNotInChannelErrorNotOnChannelErrorUserOnChannelErrorNoLoginErrorSummonDisabledErrorUsersDisabledErrorNotRegisteredErrorNeedMoreParamsErrorAlreadyRegisteredErrorNoPermForHostErrorPasswordMismatchErrorYoureBannedErrorYouWillBeBannedErrorKeySetErrorChannelIsFullErrorUnknownModeErrorInviteOnlyChanErrorBannedFromChanErrorBadChannelKeyErrorBadChannelMaskErrorNoChanelModesErrorBanListFullErrorNoPrivilegesErrorChanOpIsNeededErrorCantKillServerErrorRestrictedErrorOriginalOpIsNeededErrorNoOperHostErrorUModeUnknownFlagErrorUsersDontMatch"

var _ReplyCode_map = map[ReplyCode]string{
	1:   _ReplyCode_name[0:12],
	2:   _ReplyCode_name[12:25],
	3:   _ReplyCode_name[25:37],
	4:   _ReplyCode_name[37:48],
	5:   _ReplyCode_name[48:59],
	200: _ReplyCode_name[59:73],
	201: _ReplyCode_name[73:93],
	202: _ReplyCode_name[93:112],
	203: _ReplyCode_name[112:129],
	204: _ReplyCode_name[129:147],
	205: _ReplyCode_name[147:161],
	206: _ReplyCode_name[161:177],
	207: _ReplyCode_name[177:194],
	208: _ReplyCode_name[194:211],
	209: _ReplyCode_name[211:226],
	210: _ReplyCode_name[226:245],
	211: _ReplyCode_name[245:263],
	212: _ReplyCode_name[263:281],
	219: _ReplyCode_name[281:296],
	221: _ReplyCode_name[296:311],
	234: _ReplyCode_name[311:326],
	235: _ReplyCode_name[326:344],
	242: _ReplyCode_name[344:360],
	243: _ReplyCode_name[360:375],
	251: _ReplyCode_name[375:391],
	252: _ReplyCode_name[391:403],
	253: _ReplyCode_name[403:420],
	254: _ReplyCode_name[420:438],
	255: _ReplyCode_name[438:450],
	256: _ReplyCode_name[450:462],
	257: _ReplyCode_name[462:476],
	258: _ReplyCode_name[476:490],
	259: _ReplyCode_name[490:505],
	261: _ReplyCode_name[505:518],
	262: _ReplyCode_name[518:531],
	263: _ReplyCode_name[531:544],
	301: _ReplyCode_name[544:553],
	302: _ReplyCode_name[553:566],
	303: _ReplyCode_name[566:575],
	305: _ReplyCode_name[575:586],
	306: _ReplyCode_name[586:598],
	311: _ReplyCode_name[598:612],
	312: _ReplyCode_name[612:628],
	313: _ReplyCode_name[628:646],
	314: _ReplyCode_name[646:661],
	315: _ReplyCode_name[661:674],
	317: _ReplyCode_name[674:688],
	318: _ReplyCode_name[688:703],
	319: _ReplyCode_name[703:721],
	321: _ReplyCode_name[721:735],
	322: _ReplyCode_name[735:744],
	323: _ReplyCode_name[744:756],
	324: _ReplyCode_name[756:774],
	325: _ReplyCode_name[774:789],
	331: _ReplyCode_name[789:801],
	332: _ReplyCode_name[801:811],
	341: _ReplyCode_name[811:824],
	342: _ReplyCode_name[824:838],
	346: _ReplyCode_name[838:853],
	347: _ReplyCode_name[853:873],
	348: _ReplyCode_name[873:888],
	349: _ReplyCode_name[888:908],
	351: _ReplyCode_name[908:920],
	352: _ReplyCode_name[920:933],
	353: _ReplyCode_name[933:947],
	364: _ReplyCode_name[947:957],
	365: _ReplyCode_name[957:972],
	366: _ReplyCode_name[972:987],
	367: _ReplyCode_name[987:999],
	368: _ReplyCode_name[999:1016],
	369: _ReplyCode_name[1016:1032],
	371: _ReplyCode_name[1032:1041],
	372: _ReplyCode_name[1041:1050],
	374: _ReplyCode_name[1050:1064],
	375: _ReplyCode_name[1064:1078],
	376: _ReplyCode_name[1078:1092],
	381: _ReplyCode_name[1092:1106],
	382: _ReplyCode_name[1106:1120],
	383: _ReplyCode_name[1120:1137],
	391: _ReplyCode_name[1137:1146],
	392: _ReplyCode_name[1146:1160],
	393: _ReplyCode_name[1160:1170],
	394: _ReplyCode_name[1170:1185],
	395: _ReplyCode_name[1185:1197],
	401: _ReplyCode_name[1197:1216],
	402: _ReplyCode_name[1216:1233],
	403: _ReplyCode_name[1233:1251],
	404: _ReplyCode_name[1251:1275],
	405: _ReplyCode_name[1275:1295],
	406: _ReplyCode_name[1295:1317],
	407: _ReplyCode_name[1317:1336],
	408: _ReplyCode_name[1336:1354],
	409: _ReplyCode_name[1354:1367],
	411: _ReplyCode_name[1367:1383],
	412: _ReplyCode_name[1383:1400],
	413: _ReplyCode_name[1400:1415],
	414: _ReplyCode_name[1415:1432],
	415: _ReplyCode_name[1432:1444],
	421: _ReplyCode_name[1444:1463],
	422: _ReplyCode_name[1463:1474],
	423: _ReplyCode_name[1474:1490],
	424: _ReplyCode_name[1490:1504],
	431: _ReplyCode_name[1504:1524],
	432: _ReplyCode_name[1524:1546],
	433: _ReplyCode_name[1546:1564],
	436: _ReplyCode_name[1564:1586],
	437: _ReplyCode_name[1586:1610],
	441: _ReplyCode_name[1610:1631],
	442: _ReplyCode_name[1631:1648],
	443: _ReplyCode_name[1648:1666],
	444: _ReplyCode_name[1666:1678],
	445: _ReplyCode_name[1678:1697],
	446: _ReplyCode_name[1697:1715],
	451: _ReplyCode_name[1715:1733],
	461: _ReplyCode_name[1733:1752],
	462: _ReplyCode_name[1752:1774],
	463: _ReplyCode_name[1774:1792],
	464: _ReplyCode_name[1792:1813],
	465: _ReplyCode_name[1813:1829],
	466: _ReplyCode_name[1829:1849],
	467: _ReplyCode_name[1849:1860],
	471: _ReplyCode_name[1860:1878],
	472: _ReplyCode_name[1878:1894],
	473: _ReplyCode_name[1894:1913],
	474: _ReplyCode_name[1913:1932],
	475: _ReplyCode_name[1932:1950],
	476: _ReplyCode_name[1950:1969],
	477: _ReplyCode_name[1969:1987],
	478: _ReplyCode_name[1987:2003],
	481: _ReplyCode_name[2003:2020],
	482: _ReplyCode_name[2020:2039],
	483: _ReplyCode_name[2039:2058],
	484: _ReplyCode_name[2058:2073],
	485: _ReplyCode_name[2073:2096],
	491: _ReplyCode_name[2096:2111],
	501: _ReplyCode_name[2111:2132],
	502: _ReplyCode_name[2132:2151],
}

func (i ReplyCode) String() string {
	if str, ok := _ReplyCode_map[i]; ok {
		return str
	}
	return fmt.Sprintf("ReplyCode(%d)", i)
}
