package irc

import (
	"fmt"
	"strconv"
)

//go:generate stringer -type=ReplyCode

type ReplyCode int

// Client-Server
const (
	ReplyWelcome  ReplyCode = 001
	ReplyYourHost ReplyCode = 002
	ReplyCreated  ReplyCode = 003
	ReplyMyInfo   ReplyCode = 004
	ReplyBounce   ReplyCode = 005
)

// Command responses
const (
	ReplyTraceLink       ReplyCode = 200
	ReplyTraceConnecting ReplyCode = 201
	ReplyTraceHandshake  ReplyCode = 202
	ReplyTraceUnknown    ReplyCode = 203
	ReplyTraceOperator   ReplyCode = 204
	ReplyTraceUser       ReplyCode = 205
	ReplyTraceServer     ReplyCode = 206
	ReplyTraceService    ReplyCode = 207
	ReplyTraceNewType    ReplyCode = 208
	ReplyTraceClass      ReplyCode = 209
	ReplyTraceReconnect  ReplyCode = 210
	ReplyStatsLinkInfo   ReplyCode = 211
	ReplyStatsCommands   ReplyCode = 212
	ReplyEndOfStats      ReplyCode = 219
	ReplyUserModeIs      ReplyCode = 221
	ReplyServerList      ReplyCode = 234
	ReplyServerListEnd   ReplyCode = 235
	ReplyStatsUpTime     ReplyCode = 242
	ReplyStatsOLine      ReplyCode = 243
	ReplyLUserClient     ReplyCode = 251
	ReplyLUserOp         ReplyCode = 252
	ReplyLUserUnknown    ReplyCode = 253
	ReplyLUserChannels   ReplyCode = 254
	ReplyLUserMe         ReplyCode = 255
	ReplyAdminMe         ReplyCode = 256
	ReplyAdminLoc1       ReplyCode = 257
	ReplyAdminLoc2       ReplyCode = 258
	ReplyAdminEmail      ReplyCode = 259
	ReplyTraceLog        ReplyCode = 261
	ReplyTraceEnd        ReplyCode = 262
	ReplyTryAgain        ReplyCode = 263
	ReplyAway            ReplyCode = 301
	ReplyUserHost        ReplyCode = 302
	ReplyISON            ReplyCode = 303
	ReplyUnAway          ReplyCode = 305
	ReplyNowAway         ReplyCode = 306
	ReplyWhoIsUser       ReplyCode = 311
	ReplyWhoIsServer     ReplyCode = 312
	ReplyWhoIsOperator   ReplyCode = 313
	ReplyWhoWasUser      ReplyCode = 314
	ReplyEndOfWho        ReplyCode = 315
	ReplyWhoIsIdle       ReplyCode = 317
	ReplyEndOfWhoIs      ReplyCode = 318
	ReplyWhoIsChannels   ReplyCode = 319
	ReplyListStart       ReplyCode = 321
	ReplyList            ReplyCode = 322
	ReplyListEnd         ReplyCode = 323
	ReplyChannelModeIs   ReplyCode = 324
	ReplyUniqueOpIs      ReplyCode = 325
	ReplyNoTopic         ReplyCode = 331
	ReplyTopic           ReplyCode = 332
	ReplyInviting        ReplyCode = 341
	ReplySummoning       ReplyCode = 342
	ReplyInviteList      ReplyCode = 346
	ReplyEndOfInviteList ReplyCode = 347
	ReplyExceptList      ReplyCode = 348
	ReplyEndOfExceptList ReplyCode = 349
	ReplyVersion         ReplyCode = 351
	ReplyWhoReply        ReplyCode = 352
	ReplyNameReply       ReplyCode = 353
	ReplyLinks           ReplyCode = 364
	ReplyEndOfLinks      ReplyCode = 365
	ReplyEndOfNames      ReplyCode = 366
	ReplyBanList         ReplyCode = 367
	ReplyEndOfBanList    ReplyCode = 368
	ReplyEndOfWhoWas     ReplyCode = 369
	ReplyInfo            ReplyCode = 371
	ReplyMOTD            ReplyCode = 372
	ReplyEndOfInfo       ReplyCode = 374
	ReplyMOTDStart       ReplyCode = 375
	ReplyEndOfMOTD       ReplyCode = 376
	ReplyYoureOper       ReplyCode = 381
	ReplyRehashing       ReplyCode = 382
	ReplyYoureService    ReplyCode = 383
	ReplyTime            ReplyCode = 391
	ReplyUserStart       ReplyCode = 392
	ReplyUsers           ReplyCode = 393
	ReplyEndOfUsers      ReplyCode = 394
	ReplyNoUsers         ReplyCode = 395
)

// Error responses
const (
	ErrorNoSuchNickname      ReplyCode = 401
	ErrorNoSuchServer        ReplyCode = 402
	ErrorNoSuchChannel       ReplyCode = 403
	ErrorCannotSendToChannel ReplyCode = 404
	ErrorTooManyChannels     ReplyCode = 405
	ErrorWasNoSuchNickname   ReplyCode = 406
	ErrorTooManyTargets      ReplyCode = 407
	ErrorNoSuchService       ReplyCode = 408
	ErrorNoOrigin            ReplyCode = 409
	ErrorNoRecipient         ReplyCode = 411
	ErrorNoTextToSend        ReplyCode = 412
	ErrorNoTopLevel          ReplyCode = 413
	ErrorWildTopLevel        ReplyCode = 414
	ErrorBadMask             ReplyCode = 415
	ErrorUnknownCommand      ReplyCode = 421
	ErrorNoMOTD              ReplyCode = 422
	ErrorNoAdminInfo         ReplyCode = 423
	ErrorFileError           ReplyCode = 424
	ErrorNoNicknameGiven     ReplyCode = 431
	ErrorErroneousNickname   ReplyCode = 432
	ErrorNicknameInUse       ReplyCode = 433
	ErrorNicknameCollision   ReplyCode = 436
	ErrorUnavailableResource ReplyCode = 437
	ErrorUserNotInChannel    ReplyCode = 441
	ErrorNotOnChannel        ReplyCode = 442
	ErrorUserOnChannel       ReplyCode = 443
	ErrorNoLogin             ReplyCode = 444
	ErrorSummonDisabled      ReplyCode = 445
	ErrorUsersDisabled       ReplyCode = 446
	ErrorNotRegistered       ReplyCode = 451
	ErrorNeedMoreParams      ReplyCode = 461
	ErrorAlreadyRegistered   ReplyCode = 462
	ErrorNoPermForHost       ReplyCode = 463
	ErrorPasswordMismatch    ReplyCode = 464
	ErrorYoureBanned         ReplyCode = 465
	ErrorYouWillBeBanned     ReplyCode = 466
	ErrorKeySet              ReplyCode = 467
	ErrorChannelIsFull       ReplyCode = 471
	ErrorUnknownMode         ReplyCode = 472
	ErrorInviteOnlyChan      ReplyCode = 473
	ErrorBannedFromChan      ReplyCode = 474
	ErrorBadChannelKey       ReplyCode = 475
	ErrorBadChannelMask      ReplyCode = 476
	ErrorNoChanelModes       ReplyCode = 477
	ErrorBanListFull         ReplyCode = 478
	ErrorNoPrivileges        ReplyCode = 481
	ErrorChanOpIsNeeded      ReplyCode = 482
	ErrorCantKillServer      ReplyCode = 483
	ErrorRestricted          ReplyCode = 484
	ErrorOriginalOpIsNeeded  ReplyCode = 485
	ErrorNoOperHost          ReplyCode = 491
	ErrorUModeUnknownFlag    ReplyCode = 501
	ErrorUsersDontMatch      ReplyCode = 502
)

type Reply struct {
	unhandledMessage
}

func NewReply(code ReplyCode, args ...string) Reply {
	return Reply{
		unhandledMessage{
			command:   fmt.Sprintf("%03d", code),
			arguments: args,
		},
	}
}

func (r Reply) Code() ReplyCode {
	i, err := strconv.ParseInt(r.command, 10, 32)
	if err != nil {
		panic(err) // Shouldn't ever happen
	}
	return ReplyCode(i)
}
