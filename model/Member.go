package model

import "time"

type Member struct {
	ID                uint      `json:"ID" gorm:"column:ID;primary_key"`
	RefID             uint      `json:"refId" gorm:"column:RefID_FK;not null;index:RefID_FK"`
	RefUsername       string    `json:"refUsername" gorm:"column:RefUsername;type:varchar(60);not null"`
	RefIDTreeLevel    int       `json:"refIdTreeLevel" gorm:"column:RefIDTreeLevel;not null;index:RefIDTreeLevel"`
	ImageCoverID      uint      `json:"imageCoverID" gorm:"column:ImageCoverID_FK;not null"`
	Username          string    `json:"username" gorm:"column:Username;type:varchar(100);not null;index:Username"`
	Password          string    `json:"password" gorm:"column:Password;type:varchar(100);not null"`
	SecPassword       string    `json:"secPassword" gorm:"column:SecPassword;type:varchar(70);not null"`
	PasswordEncode    string    `json:"passwordEncode" gorm:"column:PasswordEncode;type:varchar(60);not null"`
	SecPasswordEncode string    `json:"secPasswordEncode" gorm:"column:SecPasswordEncode;type:varchar(70);not null"`
	Name              string    `json:"name" gorm:"column:Name;type:varchar(100);not null"`
	Email             string    `json:"email" gorm:"column:Email;type:varchar(100);unique_index;not null"`
	Phone             string    `json:"phone" gorm:"column:Phone;type:varchar(40);not null"`
	Code              string    `json:"code" gorm:"column:Code;type:varchar(12);not null;index:Code"`
	Lang              string    `json:"lang" gorm:"column:Lang;type:varchar(10);not null DEFAULT 'en'"`
	Country           string    `json:"country" gorm:"column:Country;type:varchar(30);not null"`
	UWallet           float64   `json:"uWallet" gorm:"column:UWallet;type:decimal(20,6);not null;default:0.00"`
	CWallet           float64   `json:"cWallet" gorm:"column:CWallet;type:decimal(20,6);not null;default:0.00"`
	RWallet           float64   `json:"rWallet" gorm:"column:RWallet;type:decimal(20,6);not null;default:0.00"`
	IWallet           float64   `json:"iWallet" gorm:"column:IWallet;type:decimal(20,6);not null;default:0.00"`
	WWallet           float64   `json:"wWallet" gorm:"column:WWallet;type:decimal(20,6);not null;default:0.00"`
	SWallet           float64   `json:"sWallet" gorm:"column:SWallet;type:decimal(20,6);not null;default:0.00"`
	MWallet           float64   `json:"mWallet" gorm:"column:MWallet;type:decimal(20,8);not null;default:0.00"`
	OWallet           float64   `json:"oWallet" gorm:"column:OWallet;type:decimal(20,6);not null;default:0.00"`
	JoinedDate        time.Time `json:"joinedDate" gorm:"column:JoinedDate;type:datetime;not null"`
	Status            string    `json:"status" gorm:"column:Status;type:varchar(30);not null DEFAULT 'demo';index:Status"`
	Ranking           int8      `json:"ranking" gorm:"column:Ranking;type:tinyint(4);not null"`
	Remark            string    `json:"remark" gorm:"column:Remark;type:text;not null"`
	TotalDownlines    int       `json:"totalDownlines" gorm:"column:TotalDownlines;type:int(11);not null"`
	LastLogin         time.Time `json:"lastLogin" gorm:"column:LastLogin;type:datetime;not null"`
	LastLogin2        time.Time `json:"lastLogin2" gorm:"column:LastLogin2;type:datetime;not null"`
	IsRobot           int8      `json:"isRobot" gorm:"column:IsRobot;type:tinyint(4);not null DEFAULT 0;index:IsRobot"`
	IsVip             int8      `json:"isVip" gorm:"column:IsVip;type:tinyint(4);not null DEFAULT 0;index:IsVip"`
	VIPDaily          int       `json:"vipDaily" gorm:"column:VIPDaily;type:int(11);not null DEFAULT 0"`
	WalletAddress     string    `json:"walletAddress" gorm:"column:WalletAddress;type:varchar(100);not null"`
	AutoRanking       string    `json:"autoRanking" gorm:"column:AutoRanking;type:enum('y','n');not null DEFAULT 'y'"`
	TransferTo        uint      `json:"transferTo" gorm:"column:TransferTo;type:int(11);not null;index:TransferTo"`
	MaxWithdraw       float64   `json:"maxWithdraw" gorm:"column:MaxWithdraw;type:decimal(12,2);not null"`
	CreatedAt         time.Time `json:"createdAt" gorm:"column:CreatedAt;type:datetime;not null"`
	UpdatedAt         time.Time `json:"updatedAt" gorm:"column:UpdatedAt;type:datetime;not null COMMENT 'Modified Date'"`
}

func (Member) TableName() string {
	return "jk_members"
}
