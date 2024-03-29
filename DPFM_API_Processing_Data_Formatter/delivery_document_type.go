package dpfm_dpfm_api_processing_data_formatter

type DeliveryDocumentSDC struct {
	HeaderAndItem    DeliveryDocumentHeaderAndItem    `json:"HeaderAndItem"`
	HeaderAndPartner DeliveryDocumentHeaderAndPartner `json:"HeaderAndPartner"`
	HeaderAndAddress DeliveryDocumentHeaderAndAddress `json:"HeaderAndAddress"`
}

type DeliveryDocumentHeaderAndItem struct {
	Header DeliveryDocumentHeader `json:"Header"`
	Item   []DeliveryDocumentItem `json:"Item"`
}

type DeliveryDocumentHeaderAndPartner struct {
	Header  DeliveryDocumentHeader    `json:"Header"`
	Partner []DeliveryDocumentPartner `json:"Partner"`
}

type DeliveryDocumentHeaderAndAddress struct {
	Header  DeliveryDocumentHeader    `json:"Header"`
	Address []DeliveryDocumentAddress `json:"Address"`
}

type DeliveryDocumentHeader struct {
	DeliveryDocument                       int                       `json:"DeliveryDocument"`
	SupplyChainRelationshipID              *int                      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      *int                      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID *int                      `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipBillingID       *int                      `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID       *int                      `json:"SupplyChainRelationshipPaymentID"`
	Buyer                                  *int                      `json:"Buyer"`
	Seller                                 *int                      `json:"Seller"`
	DeliverToParty                         *int                      `json:"DeliverToParty"`
	DeliverFromParty                       *int                      `json:"DeliverFromParty"`
	DeliverToPlant                         *string                   `json:"DeliverToPlant"`
	DeliverFromPlant                       *string                   `json:"DeliverFromPlant"`
	BillToParty                            *int                      `json:"BillToParty"`
	BillFromParty                          *int                      `json:"BillFromParty"`
	BillToCountry                          *string                   `json:"BillToCountry"`
	BillFromCountry                        *string                   `json:"BillFromCountry"`
	Payer                                  *int                      `json:"Payer"`
	Payee                                  *int                      `json:"Payee"`
	IsExportImport                         *bool                     `json:"IsExportImport"`
	DeliverToPlantTimeZone                 *string                   `json:"DeliverToPlantTimeZone"`
	DeliverFromPlantTimeZone               *string                   `json:"DeliverFromPlantTimeZone"`
	ReferenceDocument                      *int                      `json:"ReferenceDocument"`
	ReferenceDocumentItem                  *int                      `json:"ReferenceDocumentItem"`
	OrderID                                *int                      `json:"OrderID"`
	OrderItem                              *int                      `json:"OrderItem"`
	ContractType                           *string                   `json:"ContractType"`
	OrderValidityStartDate                 *string                   `json:"OrderValidityStartDate"`
	OrderValidityEndDate                   *string                   `json:"OrderValidityEndDate"`
	DocumentDate                           *string                   `json:"DocumentDate"`
	PlannedGoodsIssueDate                  *string                   `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime                  *string                   `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDate                *string                   `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime                *string                   `json:"PlannedGoodsReceiptTime"`
	InvoiceDocumentDate                    *string                   `json:"InvoiceDocumentDate"`
	HeaderCompleteDeliveryIsDefined        *bool                     `json:"HeaderCompleteDeliveryIsDefined"`
	HeaderDeliveryStatus                   *string                   `json:"HeaderDeliveryStatus"`
	CreationDate                           *string                   `json:"CreationDate"`
	CreationTime                           *string                   `json:"CreationTime"`
	LastChangeDate                         *string                   `json:"LastChangeDate"`
	LastChangeTime                         *string                   `json:"LastChangeTime"`
	GoodsIssueOrReceiptSlipNumber          *string                   `json:"GoodsIssueOrReceiptSlipNumber"`
	HeaderBillingStatus                    *string                   `json:"HeaderBillingStatus"`
	HeaderBillingConfStatus                *string                   `json:"HeaderBillingConfStatus"`
	HeaderBillingBlockStatus               *bool                     `json:"HeaderBillingBlockStatus"`
	HeaderGrossWeight                      *float32                  `json:"HeaderGrossWeight"`
	HeaderNetWeight                        *float32                  `json:"HeaderNetWeight"`
	HeaderWeightUnit                       *string                   `json:"HeaderWeightUnit"`
	Incoterms                              *string                   `json:"Incoterms"`
	TransactionCurrency                    *string                   `json:"TransactionCurrency"`
	HeaderDeliveryBlockStatus              *bool                     `json:"HeaderDeliveryBlockStatus"`
	HeaderIssuingBlockStatus               *bool                     `json:"HeaderIssuingBlockStatus"`
	HeaderReceivingBlockStatus             *bool                     `json:"HeaderReceivingBlockStatus"`
	IsCancelled                            *bool                     `json:"IsCancelled"`
	IsMarkedForDeletion                    *bool                     `json:"IsMarkedForDeletion"`
	Item                                   []DeliveryDocumentItem    `json:"DeliveryDocumentItem"`
	Partner                                []DeliveryDocumentPartner `json:"DeliveryDocumentPartner"`
	Address                                []DeliveryDocumentAddress `json:"DeliveryDocumentAddress"`
}

type DeliveryDocumentItem struct {
	DeliveryDocument                              int      `json:"DeliveryDocument"`
	DeliveryDocumentItem                          int      `json:"DeliveryDocumentItem"`
	DeliveryDocumentItemCategory                  *string  `json:"DeliveryDocumentItemCategory"`
	SupplyChainRelationshipID                     *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID             *int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID        *int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipStockConfPlantID       *int     `json:"SupplyChainRelationshipStockConfPlantID"`
	SupplyChainRelationshipProductionPlantID      *int     `json:"SupplyChainRelationshipProductionPlantID"`
	SupplyChainRelationshipBillingID              *int     `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID              *int     `json:"SupplyChainRelationshipPaymentID"`
	Buyer                                         *int     `json:"Buyer"`
	Seller                                        *int     `json:"Seller"`
	DeliverToParty                                *int     `json:"DeliverToParty"`
	DeliverFromParty                              *int     `json:"DeliverFromParty"`
	DeliverToPlant                                *string  `json:"DeliverToPlant"`
	DeliverFromPlant                              *string  `json:"DeliverFromPlant"`
	BillToParty                                   *int     `json:"BillToParty"`
	BillFromParty                                 *int     `json:"BillFromParty"`
	BillToCountry                                 *string  `json:"BillToCountry"`
	BillFromCountry                               *string  `json:"BillFromCountry"`
	Payer                                         *int     `json:"Payer"`
	Payee                                         *int     `json:"Payee"`
	DeliverToPlantStorageLocation                 *string  `json:"DeliverToPlantStorageLocation"`
	ProductIsBatchManagedInDeliverToPlant         *bool    `json:"ProductIsBatchManagedInDeliverToPlant"`
	BatchMgmtPolicyInDeliverToPlant               *string  `json:"BatchMgmtPolicyInDeliverToPlant"`
	DeliverToPlantBatch                           *string  `json:"DeliverToPlantBatch"`
	DeliverToPlantBatchValidityStartDate          *string  `json:"DeliverToPlantBatchValidityStartDate"`
	DeliverToPlantBatchValidityStartTime          *string  `json:"DeliverToPlantBatchValidityStartTime"`
	DeliverToPlantBatchValidityEndDate            *string  `json:"DeliverToPlantBatchValidityEndDate"`
	DeliverToPlantBatchValidityEndTime            *string  `json:"DeliverToPlantBatchValidityEndTime"`
	DeliverFromPlantStorageLocation               *string  `json:"DeliverFromPlantStorageLocation"`
	ProductIsBatchManagedInDeliverFromPlant       *bool    `json:"ProductIsBatchManagedInDeliverFromPlant"`
	BatchMgmtPolicyInDeliverFromPlant             *string  `json:"BatchMgmtPolicyInDeliverFromPlant"`
	DeliverFromPlantBatch                         *string  `json:"DeliverFromPlantBatch"`
	DeliverFromPlantBatchValidityStartDate        *string  `json:"DeliverFromPlantBatchValidityStartDate"`
	DeliverFromPlantBatchValidityStartTime        *string  `json:"DeliverFromPlantBatchValidityStartTime"`
	DeliverFromPlantBatchValidityEndDate          *string  `json:"DeliverFromPlantBatchValidityEndDate"`
	DeliverFromPlantBatchValidityEndTime          *string  `json:"DeliverFromPlantBatchValidityEndTime"`
	StockConfirmationBusinessPartner              *int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                        *string  `json:"StockConfirmationPlant"`
	ProductIsBatchManagedInStockConfirmationPlant *bool    `json:"ProductIsBatchManagedInStockConfirmationPlant"`
	BatchMgmtPolicyInStockConfirmationPlant       *string  `json:"BatchMgmtPolicyInStockConfirmationPlant"`
	StockConfirmationPlantBatch                   *string  `json:"StockConfirmationPlantBatch"`
	StockConfirmationPlantBatchValidityStartDate  *string  `json:"StockConfirmationPlantBatchValidityStartDate"`
	StockConfirmationPlantBatchValidityStartTime  *string  `json:"StockConfirmationPlantBatchValidityStartTime"`
	StockConfirmationPlantBatchValidityEndDate    *string  `json:"StockConfirmationPlantBatchValidityEndDate"`
	StockConfirmationPlantBatchValidityEndTime    *string  `json:"StockConfirmationPlantBatchValidityEndTime"`
	StockConfirmationPolicy                       *string  `json:"StockConfirmationPolicy"`
	StockConfirmationStatus                       *string  `json:"StockConfirmationStatus"`
	ProductionPlantBusinessPartner                *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                               *string  `json:"ProductionPlant"`
	ProductionPlantStorageLocation                *string  `json:"ProductionPlantStorageLocation"`
	ProductIsBatchManagedInProductionPlant        *bool    `json:"ProductIsBatchManagedInProductionPlant"`
	BatchMgmtPolicyInProductionPlant              *string  `json:"BatchMgmtPolicyInProductionPlant"`
	ProductionPlantBatch                          *string  `json:"ProductionPlantBatch"`
	ProductionPlantBatchValidityStartDate         *string  `json:"ProductionPlantBatchValidityStartDate"`
	ProductionPlantBatchValidityStartTime         *string  `json:"ProductionPlantBatchValidityStartTime"`
	ProductionPlantBatchValidityEndDate           *string  `json:"ProductionPlantBatchValidityEndDate"`
	ProductionPlantBatchValidityEndTime           *string  `json:"ProductionPlantBatchValidityEndTime"`
	InspectionPlan                                *int     `json:"InspectionPlan"`
	InspectionPlant                               *string  `json:"InspectionPlant"`
	InspectionOrder                               *int     `json:"InspectionOrder"`
	DeliveryDocumentItemText                      *string  `json:"DeliveryDocumentItemText"`
	DeliveryDocumentItemTextByBuyer               *string  `json:"DeliveryDocumentItemTextByBuyer"`
	DeliveryDocumentItemTextBySeller              *string  `json:"DeliveryDocumentItemTextBySeller"`
	Product                                       *string  `json:"Product"`
	ProductStandardID                             *string  `json:"ProductStandardID"`
	ProductGroup                                  *string  `json:"ProductGroup"`
	BaseUnit                                      *string  `json:"BaseUnit"`
	OriginalQuantityInBaseUnit                    *float32 `json:"OriginalQuantityInBaseUnit"`
	DeliveryUnit                                  *string  `json:"DeliveryUnit"`
	PlannedGoodsIssueDate                         *string  `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime                         *string  `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDate                       *string  `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime                       *string  `json:"PlannedGoodsReceiptTime"`
	PlannedGoodsIssueQuantity                     *float32 `json:"PlannedGoodsIssueQuantity"`
	PlannedGoodsIssueQtyInBaseUnit                *float32 `json:"PlannedGoodsIssueQtyInBaseUnit"`
	PlannedGoodsReceiptQuantity                   *float32 `json:"PlannedGoodsReceiptQuantity"`
	PlannedGoodsReceiptQtyInBaseUnit              *float32 `json:"PlannedGoodsReceiptQtyInBaseUnit"`
	ActualGoodsIssueDate                          *string  `json:"ActualGoodsIssueDate"`
	ActualGoodsIssueTime                          *string  `json:"ActualGoodsIssueTime"`
	ActualGoodsReceiptDate                        *string  `json:"ActualGoodsReceiptDate"`
	ActualGoodsReceiptTime                        *string  `json:"ActualGoodsReceiptTime"`
	ActualGoodsIssueQuantity                      *float32 `json:"ActualGoodsIssueQuantity"`
	ActualGoodsIssueQtyInBaseUnit                 *float32 `json:"ActualGoodsIssueQtyInBaseUnit"`
	ActualGoodsReceiptQuantity                    *float32 `json:"ActualGoodsReceiptQuantity"`
	ActualGoodsReceiptQtyInBaseUnit               *float32 `json:"ActualGoodsReceiptQtyInBaseUnit"`
	CreationDate                                  *string  `json:"CreationDate"`
	CreationTime                                  *string  `json:"CreationTime"`
	LastChangeDate                                *string  `json:"LastChangeDate"`
	LastChangeTime                                *string  `json:"LastChangeTime"`
	ItemBillingStatus                             *string  `json:"ItemBillingStatus"`
	ItemCompleteDeliveryIsDefined                 *bool    `json:"ItemCompleteDeliveryIsDefined"`
	ItemGrossWeight                               *float32 `json:"ItemGrossWeight"`
	ItemNetWeight                                 *float32 `json:"ItemNetWeight"`
	ItemWeightUnit                                *string  `json:"ItemWeightUnit"`
	InternalCapacityQuantity                      *float32 `json:"InternalCapacityQuantity"`
	InternalCapacityQuantityUnit                  *string  `json:"InternalCapacityQuantityUnit"`
	ItemIsBillingRelevant                         *bool    `json:"ItemIsBillingRelevant"`
	NetAmount                                     *float32 `json:"NetAmount"`
	TaxAmount                                     *float32 `json:"TaxAmount"`
	GrossAmount                                   *float32 `json:"GrossAmount"`
	OrderID                                       *int     `json:"OrderID"`
	OrderItem                                     *int     `json:"OrderItem"`
	OrderType                                     *string  `json:"OrderType"`
	ContractType                                  *string  `json:"ContractType"`
	OrderValidityStartDate                        *string  `json:"OrderValidityStartDate"`
	OrderValidityEndDate                          *string  `json:"OrderValidityEndDate"`
	PaymentTerms                                  *string  `json:"PaymentTerms"`
	DueCalculationBaseDate                        *string  `json:"DueCalculationBaseDate"`
	PaymentDueDate                                *string  `json:"PaymentDueDate"`
	NetPaymentDays                                *int     `json:"NetPaymentDays"`
	PaymentMethod                                 *string  `json:"PaymentMethod"`
	InvoicePeriodStartDate                        *string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate                          *string  `json:"InvoicePeriodEndDate"`
	ConfirmedDeliveryDate                         *string  `json:"ConfirmedDeliveryDate"`
	Project                                       *string  `json:"Project"`
	ReferenceDocument                             *int     `json:"ReferenceDocument"`
	ReferenceDocumentItem                         *int     `json:"ReferenceDocumentItem"`
	TransactionTaxClassification                  *string  `json:"TransactionTaxClassification"`
	ProductTaxClassificationBillToCountry         *string  `json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry       *string  `json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification                      *string  `json:"DefinedTaxClassification"`
	AccountAssignmentGroup                        *string  `json:"AccountAssignmentGroup"`
	ProductAccountAssignmentGroup                 *string  `json:"ProductAccountAssignmentGroup"`
	TaxCode                                       *string  `json:"TaxCode"`
	TaxRate                                       *float32 `json:"TaxRate"`
	CountryOfOrigin                               *string  `json:"CountryOfOrigin"`
	CountryOfOriginLanguage                       *string  `json:"CountryOfOriginLanguage"`
	ItemDeliveryBlockStatus                       *bool    `json:"ItemDeliveryBlockStatus"`
	ItemIssuingBlockStatus                        *bool    `json:"ItemIssuingBlockStatus"`
	ItemReceivingBlockStatus                      *bool    `json:"ItemReceivingBlockStatus"`
	ItemBillingBlockStatus                        *bool    `json:"ItemBillingBlockStatus"`
	IsCancelled                                   *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                           *bool    `json:"IsMarkedForDeletion"`
}

type DeliveryDocumentAddress struct {
	DeliveryDocument int     `json:"DeliveryDocument"`
	AddressID        int     `json:"AddressID"`
	PostalCode       *string `json:"PostalCode"`
	LocalRegion      *string `json:"LocalRegion"`
	Country          *string `json:"Country"`
	District         *string `json:"District"`
	StreetName       *string `json:"StreetName"`
	CityName         *string `json:"CityName"`
	Building         *string `json:"Building"`
	Floor            *int    `json:"Floor"`
	Room             *int    `json:"Room"`
}

type DeliveryDocumentPartner struct {
	DeliveryDocument        int     `json:"DeliveryDocument"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
	AddressID               *int    `json:"AddressID"`
}
