package model

import "time"

// CREATE TABLE `jk_members` (
//   `ID` bigint(20) NOT NULL,
//   `RefID_FK` int(11) NOT NULL,
//   `RefUsername` varchar(60) NOT NULL,
//   `RefIDTreeLevel` int(11) NOT NULL,
//   `ImageCoverID_FK` int(11) NOT NULL,
//   `Username` varchar(100) NOT NULL,
//   `Password` varchar(100) NOT NULL,
//   `SecPassword` varchar(70) NOT NULL,
//   `PasswordEncode` varchar(60) NOT NULL,
//   `SecPasswordEncode` varchar(70) NOT NULL,
//   `Name` varchar(100) NOT NULL,
//   `Email` varchar(100) NOT NULL,
//   `Phone` varchar(40) NOT NULL,
//   `Code` varchar(12) NOT NULL,
//   `Lang` varchar(10) NOT NULL DEFAULT 'en',
//   `Country` varchar(30) NOT NULL,
//   `UWallet` decimal(20,6) NOT NULL COMMENT 'USDT',
//   `CWallet` decimal(20,6) NOT NULL COMMENT 'CashChip',
//   `RWallet` decimal(20,6) NOT NULL COMMENT 'RollingChip',
//   `IWallet` decimal(20,6) NOT NULL COMMENT 'Insurance',
//   `WWallet` decimal(20,6) NOT NULL COMMENT 'WINT',
//   `SWallet` decimal(20,6) NOT NULL COMMENT 'Reward',
//   `MWallet` decimal(20,8) NOT NULL,
//   `OWallet` decimal(20,6) NOT NULL,
//   `JoinedDate` datetime NOT NULL,
//   `Status` varchar(30) NOT NULL DEFAULT 'demo' COMMENT 'active, terminated, demo',
//   `Ranking` tinyint(4) NOT NULL,
//   `Remark` text NOT NULL,
//   `TotalDownlines` int(11) NOT NULL,
//   `LastLogin` datetime NOT NULL,
//   `LastLogin2` datetime NOT NULL,
//   `IsRobot` tinyint(4) NOT NULL DEFAULT 0,
//   `IsVip` tinyint(4) NOT NULL DEFAULT 0,
//   `VIPDaily` int(11) NOT NULL DEFAULT 0,
//   `WalletAddress` varchar(100) NOT NULL,
//   `AutoRanking` enum('y','n') NOT NULL DEFAULT 'y',
//   `TransferTo` int(11) NOT NULL,
//   `MaxWithdraw` decimal(12,2) NOT NULL,
//   `CreatedAt` datetime NOT NULL,
//   `UpdatedAt` datetime NOT NULL COMMENT 'Modified Date'
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

// ALTER TABLE `jk_members`
//   ADD PRIMARY KEY (`ID`),
//   ADD KEY `Username` (`Username`),
//   ADD KEY `Status` (`Status`),
//   ADD KEY `RefIDTreeLevel` (`RefIDTreeLevel`),
//   ADD KEY `RefID_FK` (`RefID_FK`),
//   ADD KEY `Code` (`Code`),
//   ADD KEY `IsRobot` (`IsRobot`),
//   ADD KEY `IsVip` (`IsVip`),
//   ADD KEY `TransferTo` (`TransferTo`);

type Member struct {
	ID                uint      `gorm:"column:ID;primary_key"`
	RefID             uint      `gorm:"column:RefID_FK;not null;index:RefID_FK"`
	RefUsername       string    `gorm:"column:RefUsername;type:varchar(60);not null"`
	RefIDTreeLevel    int       `gorm:"column:RefIDTreeLevel;not null;index:RefIDTreeLevel"`
	ImageCoverID      uint      `gorm:"column:ImageCoverID_FK;not null"`
	Username          string    `gorm:"column:Username;type:varchar(100);not null;index:Username"`
	Password          string    `gorm:"column:Password;type:varchar(100);not null"`
	SecPassword       string    `gorm:"column:SecPassword;type:varchar(70);not null"`
	PasswordEncode    string    `gorm:"column:PasswordEncode;type:varchar(60);not null"`
	SecPasswordEncode string    `gorm:"column:SecPasswordEncode;type:varchar(70);not null"`
	Name              string    `gorm:"column:Name;type:varchar(100);not null"`
	Email             string    `gorm:"column:Email;type:varchar(100);unique_index;not null"`
	Phone             string    `gorm:"column:Phone;type:varchar(40);not null"`
	Code              string    `gorm:"column:Code;type:varchar(12);not null;index:Code"`
	Lang              string    `gorm:"column:Lang;type:varchar(10);not null DEFAULT 'en'"`
	Country           string    `gorm:"column:Country;type:varchar(30);not null"`
	UWallet           float64   `gorm:"column:UWallet;type:decimal(20,6);not null;default:0.00"`
	CWallet           float64   `gorm:"column:CWallet;type:decimal(20,6);not null;default:0.00"`
	RWallet           float64   `gorm:"column:RWallet;type:decimal(20,6);not null;default:0.00"`
	IWallet           float64   `gorm:"column:IWallet;type:decimal(20,6);not null;default:0.00"`
	WWallet           float64   `gorm:"column:WWallet;type:decimal(20,6);not null;default:0.00"`
	SWallet           float64   `gorm:"column:SWallet;type:decimal(20,6);not null;default:0.00"`
	MWallet           float64   `gorm:"column:MWallet;type:decimal(20,8);not null;default:0.00"`
	OWallet           float64   `gorm:"column:OWallet;type:decimal(20,6);not null;default:0.00"`
	JoinedDate        time.Time `gorm:"column:JoinedDate;type:datetime;not null"`
	Status            string    `gorm:"column:Status;type:varchar(30);not null DEFAULT 'demo';index:Status"`
	Ranking           int8      `gorm:"column:Ranking;type:tinyint(4);not null"`
	Remark            string    `gorm:"column:Remark;type:text;not null"`
	TotalDownlines    int       `gorm:"column:TotalDownlines;type:int(11);not null"`
	LastLogin         time.Time `gorm:"column:LastLogin;type:datetime;not null"`
	LastLogin2        time.Time `gorm:"column:LastLogin2;type:datetime;not null"`
	IsRobot           int8      `gorm:"column:IsRobot;type:tinyint(4);not null DEFAULT 0;index:IsRobot"`
	IsVip             int8      `gorm:"column:IsVip;type:tinyint(4);not null DEFAULT 0;index:IsVip"`
	VIPDaily          int       `gorm:"column:VIPDaily;type:int(11);not null DEFAULT 0"`
	WalletAddress     string    `gorm:"column:WalletAddress;type:varchar(100);not null"`
	AutoRanking       string    `gorm:"column:AutoRanking;type:enum('y','n');not null DEFAULT 'y'"`
	TransferTo        uint      `gorm:"column:TransferTo;type:int(11);not null;index:TransferTo"`
	MaxWithdraw       float64   `gorm:"column:MaxWithdraw;type:decimal(12,2);not null"`
	CreatedAt         time.Time `gorm:"column:CreatedAt;type:datetime;not null"`
	UpdatedAt         time.Time `gorm:"column:UpdatedAt;type:datetime;not null COMMENT 'Modified Date'"`
}

func (Member) TableName() string {
	return "jk_members"
}
