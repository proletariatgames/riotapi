package riotapi

type QueueType int

const (
	CUSTOM                          QueueType = 0   // Custom games
	NORMAL_3x3                      QueueType = 8   // Normal 3v3 games
	NORMAL_5x5_BLIND                QueueType = 2   // Normal 5v5 Blind Pick games
	NORMAL_5x5_DRAFT                QueueType = 14  // Normal 5v5 Draft Pick games
	RANKED_SOLO_5x5                 QueueType = 4   // Ranked Solo 5v5 games
	RANKED_PREMADE_5x5              QueueType = 6   // Ranked Premade 5v5 games (Deprecated)
	RANKED_PREMADE_3x3              QueueType = 9   // Historical Ranked Premade 3v3 games (Deprecated)
	RANKED_FLEX_TT                  QueueType = 9   // Ranked Flex Twisted Treeline games
	RANKED_TEAM_3x3                 QueueType = 41  // Ranked Team 3v3 games (Deprecated)
	RANKED_TEAM_5x5                 QueueType = 42  // Ranked Team 5v5 games
	ODIN_5x5_BLIND                  QueueType = 16  // Dominion 5v5 Blind Pick games
	ODIN_5x5_DRAFT                  QueueType = 17  // Dominion 5v5 Draft Pick games
	BOT_5x5                         QueueType = 7   // Historical Summoner's Rift Coop vs AI games (Deprecated)
	BOT_ODIN_5x5                    QueueType = 25  // Dominion Coop vs AI games
	BOT_5x5_INTRO                   QueueType = 31  // Summoner's Rift Coop vs AI Intro Bot games
	BOT_5x5_BEGINNER                QueueType = 32  // Summoner's Rift Coop vs AI Beginner Bot games
	BOT_5x5_INTERMEDIATE            QueueType = 33  // Historical Summoner's Rift Coop vs AI Intermediate Bot games
	BOT_TT_3x3                      QueueType = 52  // Twisted Treeline Coop vs AI games
	GROUP_FINDER_5x5                QueueType = 61  // Team Builder games
	ARAM_5x5                        QueueType = 65  // ARAM games
	ONEFORALL_5x5                   QueueType = 70  // One for All games
	FIRSTBLOOD_1x1                  QueueType = 72  // Snowdown Showdown 1v1 games
	FIRSTBLOOD_2x2                  QueueType = 73  // Snowdown Showdown 2v2 games
	SR_6x6                          QueueType = 75  // Summoner's Rift 6x6 Hexakill games
	URF_5x5                         QueueType = 76  // Ultra Rapid Fire games
	ONEFORALL_MIRRORMODE_5x5        QueueType = 78  // One for All (Mirror mode)
	BOT_URF_5x5                     QueueType = 83  // Ultra Rapid Fire games played against AI games
	NIGHTMARE_BOT_5x5_RANK1         QueueType = 91  // Doom Bots Rank 1 games
	NIGHTMARE_BOT_5x5_RANK2         QueueType = 92  // Doom Bots Rank 2 games
	NIGHTMARE_BOT_5x5_RANK5         QueueType = 93  // Doom Bots Rank 5 games
	ASCENSION_5x5                   QueueType = 96  // Ascension games
	HEXAKILL                        QueueType = 98  // Twisted Treeline 6x6 Hexakill games
	BILGEWATER_ARAM_5x5             QueueType = 100 // Butcher's Bridge games
	KING_PORO_5x5                   QueueType = 300 // King Poro games
	COUNTER_PICK                    QueueType = 310 // Nemesis games
	BILGEWATER_5x5                  QueueType = 313 // Black Market Brawlers games
	SIEGE                           QueueType = 315 // Nexus Siege games
	DEFINITELY_NOT_DOMINION_5x5     QueueType = 317 // Definitely Not Dominion games
	ARURF_5X5                       QueueType = 318 // All Random URF games
	ARSR_5x5                        QueueType = 325 // All Random Summoner's Rift games
	TEAM_BUILDER_DRAFT_UNRANKED_5x5 QueueType = 400 // Normal 5v5 Draft Pick games
	TEAM_BUILDER_DRAFT_RANKED_5x5   QueueType = 410 // Ranked 5v5 Draft Pick games (Deprecated)
	TEAM_BUILDER_RANKED_SOLO        QueueType = 420 // Ranked Solo games from current season that use Team Builder matchmaking
	RANKED_FLEX_SR                  QueueType = 440 // Ranked Flex Summoner's Rift games
	ASSASSINATE_5x5                 QueueType = 600 // Blood Hunt Assassin games
	DARKSTAR_3x3                    QueueType = 610 // Darkstar games
)
