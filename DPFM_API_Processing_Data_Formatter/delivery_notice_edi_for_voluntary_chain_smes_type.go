package dpfm_dpfm_api_processing_data_formatter

type DeliveryNoticeEDIForVoluntaryChainSMEsSDC struct {
	HeaderAndItem DeliveryNoticeEDIForVoluntaryChainSMEsHeaderAndItem `json:"HeaderAndItem"`
}

type DeliveryNoticeEDIForVoluntaryChainSMEsHeaderAndItem struct {
	Header DeliveryNoticeEDIForVoluntaryChainSMEsHeader `json:"Header"`
	Item   []DeliveryNoticeEDIForVoluntaryChainSMEsItem `json:"Item"`
}

type DeliveryNoticeEDIForVoluntaryChainSMEsHeader struct {
	ExchangedDeliveryNoticeDocumentIdentifier                      string   `json:"ExchangedDeliveryNoticeDocumentIdentifier"`
	DeliveryNoticeDocument                                         *string  `json:"DeliveryNoticeDocument"`
	ExchangedDocumentContextSpecifiedTransactionIdentifier         *string  `json:"ExchangedDocumentContextSpecifiedTransactionIdentifier"`
	ExchangedDeliveryNoticeDocumentName                            *string  `json:"ExchangedDeliveryNoticeDocumentName"`
	ExchangeDeliveryNoticeDocumentTypeCode                         *string  `json:"ExchangeDeliveryNoticeDocumentTypeCode"`
	ExchangedDeliveryNoticeDocumentIssueDate                       *string  `json:"ExchangedDeliveryNoticeDocumentIssueDate"`
	ExchangedDeliveryNoticeDocumentPurposeCode                     *string  `json:"ExchangedDeliveryNoticeDocumentPurposeCode"`
	ExchangedDeliveryNoticeDocumentRequestedResponseTypeCode       *string  `json:"ExchangedDeliveryNoticeDocumentRequestedResponseTypeCode"`
	ExchangedDeliveryNoticeDocumentVersionIdentifier               *string  `json:"ExchangedDeliveryNoticeDocumentVersionIdentifier"`
	ExchangedDeliveryNoticeDocumentCategoryCode                    *string  `json:"ExchangedDeliveryNoticeDocumentCategoryCode"`
	ExchangedDeliveryNoticeDocumentSubtypeCode                     *string  `json:"ExchangedDeliveryNoticeDocumentSubtypeCode"`
	NoteDeliveryNoticeSubjectText                                  *string  `json:"NoteDeliveryNoticeSubjectText"`
	NoteDeliveryNoticeContentText                                  *string  `json:"NoteDeliveryNoticeContentText"`
	NoteDeliveryNoticeIdentifier                                   *string  `json:"NoteDeliveryNoticeIdentifier"`
	SpecifiedBinaryFileIdentifier                                  *string  `json:"SpecifiedBinaryFileIdentifier"`
	SpecifiedBinaryFileVersionIdentifier                           *string  `json:"SpecifiedBinaryFileVersionIdentifier"`
	SpecifiedBinaryFileNameText                                    *string  `json:"SpecifiedBinaryFileNameText"`
	SpecifiedBineryFileURIIdentifier                               *string  `json:"SpecifiedBineryFileURIIdentifier"`
	SpecifiedBineryFileMIMECode                                    *string  `json:"SpecifiedBineryFileMIMECode"`
	SpecifiedBineryFileEncodingCode                                *string  `json:"SpecifiedBineryFileEncodingCode"`
	SpecifiedBineryFileCode                                        *string  `json:"SpecifiedBineryFileCode"`
	SpecifiedBinaryFileDescriptionText                             *string  `json:"SpecifiedBinaryFileDescriptionText"`
	TradeSellerIdentifier                                          *string  `json:"TradeSellerIdentifier"`
	TradeSellerGlobalIdentifier                                    *string  `json:"TradeSellerGlobalIdentifier"`
	TradeSellerName                                                *string  `json:"TradeSellerName"`
	TradeBillFromPartyRegisteredIdentifier                         *string  `json:"TradeBillFromPartyRegisteredIdentifier"`
	TradeContactSellerIdentifier                                   *string  `json:"TradeContactSellerIdentifier"`
	TradeContactSellerPersonName                                   *string  `json:"TradeContactSellerPersonName"`
	TradeContactSellerDepartmentName                               *string  `json:"TradeContactSellerDepartmentName"`
	SellerTelephoneNumber                                          *string  `json:"SellerTelephoneNumber"`
	SellerFaxNumber                                                *string  `json:"SellerFaxNumber"`
	SellerEmailAddress                                             *string  `json:"SellerEmailAddress"`
	SellerAddressPostalCode                                        *string  `json:"SellerAddressPostalCode"`
	SellerAddress                                                  *string  `json:"SellerAddress"`
	TradeBuyerIdentifier                                           *string  `json:"TradeBuyerIdentifier"`
	TradeBuyerGlobalIdentifier                                     *string  `json:"TradeBuyerGlobalIdentifier"`
	TradeBuyerName                                                 *string  `json:"TradeBuyerName"`
	TradeContactBuyerIdentifier                                    *string  `json:"TradeContactBuyerIdentifier"`
	TradeContactBuyerPersonName                                    *string  `json:"TradeContactBuyerPersonName"`
	TradeContactBuyerDepartmentName                                *string  `json:"TradeContactBuyerDepartmentName"`
	BuyerTelephoneNumber                                           *string  `json:"BuyerTelephoneNumber"`
	BuyerFaxNumber                                                 *string  `json:"BuyerFaxNumber"`
	BuyerEmailAddress                                              *string  `json:"BuyerEmailAddress"`
	BuyerAddressPostalCode                                         *string  `json:"BuyerAddressPostalCode"`
	BuyerAddress                                                   *string  `json:"BuyerAddress"`
	ReferencedOrdersDocumentIssureAssignedIdentifier               *string  `json:"ReferencedOrdersDocumentIssureAssignedIdentifier"`
	ReferencedOrdersDocumentIssueDate                              *string  `json:"ReferencedOrdersDocumentIssueDate"`
	ReferencedOrdersDocumentRevisionIdentifier                     *string  `json:"ReferencedOrdersDocumentRevisionIdentifier"`
	ReferencedOrdersDocumentName                                   *string  `json:"ReferencedOrdersDocumentName"`
	ReferencedOrdersDocumentInformationText                        *string  `json:"ReferencedOrdersDocumentInformationText"`
	ReferencedOrdersDocumentInformationPurposeCode                 *string  `json:"ReferencedOrdersDocumentInformationPurposeCode"`
	ReferencedOrdersDocumentInformationSubtypeCode                 *string  `json:"ReferencedOrdersDocumentInformationSubtypeCode"`
	ProjectIdentifier                                              *string  `json:"ProjectIdentifier"`
	ProjectName                                                    *string  `json:"ProjectName"`
	ReferencedSalesOrderDocumentIssureAddignedIdentifier           *string  `json:"ReferencedSalesOrderDocumentIssureAddignedIdentifier"`
	ReferencedSalesOrderDocumentIssueDate                          *string  `json:"ReferencedSalesOrderDocumentIssueDate"`
	ReferencedSalesOrderDocumentRevisionIdentifier                 *string  `json:"ReferencedSalesOrderDocumentRevisionIdentifier"`
	ReferencedSalesOrderDocumentName                               *string  `json:"ReferencedSalesOrderDocumentName"`
	ReferencedSalesOrderDocumentInformationText                    *string  `json:"ReferencedSalesOrderDocumentInformationText"`
	ReferencedSalesOrderDocumentSubtypeCode                        *string  `json:"ReferencedSalesOrderDocumentSubtypeCode"`
	TradeShipToPartyIdentifier                                     *string  `json:"TradeShipToPartyIdentifier"`
	TradeShipToPartyGlobalIdentifier                               *string  `json:"TradeShipToPartyGlobalIdentifier"`
	TradeShipToPartyName                                           *string  `json:"TradeShipToPartyName"`
	TradeShipToPartyContactIdentifier                              *string  `json:"TradeShipToPartyContactIdentifier"`
	TradeShipToPartyContactPersonName                              *string  `json:"TradeShipToPartyContactPersonName"`
	TradeShipToPartyContactDepartmentName                          *string  `json:"TradeShipToPartyContactDepartmentName"`
	TradeShipToPartyContactPersonIdentifier                        *string  `json:"TradeShipToPartyContactPersonIdentifier"`
	ShipToPartyTelephoneNumber                                     *string  `json:"ShipToPartyTelephoneNumber"`
	ShipToPartyFaxNumber                                           *string  `json:"ShipToPartyFaxNumber"`
	ShipToPartyEmailAddress                                        *string  `json:"ShipToPartyEmailAddress"`
	ShipToPartyAddressPostalCode                                   *string  `json:"ShipToPartyAddressPostalCode"`
	ShipToPartyAddress                                             *string  `json:"ShipToPartyAddress"`
	LastShipToPartyIdentifier                                      *string  `json:"LastShipToPartyIdentifier"`
	TradeShipFromPartyIdentifier                                   *string  `json:"TradeShipFromPartyIdentifier"`
	TradeShipFromPartyName                                         *string  `json:"TradeShipFromPartyName"`
	TradeLogiName                                                  *string  `json:"TradeLogiName"`
	TradeLogiContactIdentifier                                     *string  `json:"TradeLogiContactIdentifier"`
	TradeLogiContactPersonName                                     *string  `json:"TradeLogiContactPersonName"`
	TradeLogiContactDepartmentName                                 *string  `json:"TradeLogiContactDepartmentName"`
	TradeLogiContactPersonIdentifier                               *string  `json:"TradeLogiContactPersonIdentifier"`
	LogiTelephoneNumber                                            *string  `json:"LogiTelephoneNumber"`
	SupplyChainEventIdentifier                                     *string  `json:"SupplyChainEventIdentifier"`
	InstructedTemperatureControlCode                               *string  `json:"InstructedTemperatureControlCode"`
	TradeTaxCalculatedAmount                                       *float32 `json:"TradeTaxCalculatedAmount"`
	TradeTaxTypeCode                                               *string  `json:"TradeTaxTypeCode"`
	TradeTaxBasisAmount                                            *float32 `json:"TradeTaxBasisAmount"`
	TradeTaxCategoryCode                                           *string  `json:"TradeTaxCategoryCode"`
	TradeTaxCategoryName                                           *string  `json:"TradeTaxCategoryName"`
	TradeTaxRatePercent                                            *float32 `json:"TradeTaxRatePercent"`
	TradeTaxGrandTotalAmount                                       *float32 `json:"TradeTaxGrandTotalAmount"`
	TradeTaxCalculationMethod                                      *string  `json:"TradeTaxCalculationMethod"`
	TradeSettlementMonetarySummationTotalTaxAmount                 *float32 `json:"TradeSettlementMonetarySummationTotalTaxAmount"`
	TradeDeliveryNoticeSettlementMonetarySummationGrandTotalAmount *float32 `json:"TradeDeliveryNoticeSettlementMonetarySummationGrandTotalAmount"`
	TradeDeliveryNoticeSettlementMonetarySummationNetTotalAmount   *float32 `json:"TradeDeliveryNoticeSettlementMonetarySummationNetTotalAmount"`
	TradeDeliveryNoticeMonetarySummationIncludingTaxesTotalAmount  *float32 `json:"TradeDeliveryNoticeMonetarySummationIncludingTaxesTotalAmount"`
}

type DeliveryNoticeEDIForVoluntaryChainSMEsItem struct {
	ExchangedDeliveryNoticeDocumentIdentifier                                      string   `json:"ExchangedDeliveryNoticeDocumentIdentifier"`
	DeliveryNoticeDocumentItemlineIdentifier                                       string   `json:"DeliveryNoticeDocumentItemlineIdentifier"`
	DeliveryNoticeDocumentItemlineStatusCode                                       *string  `json:"DeliveryNoticeDocumentItemlineStatusCode"`
	DeliveryNoticeDocumentItemlineStatusReasonCode                                 *string  `json:"DeliveryNoticeDocumentItemlineStatusReasonCode"`
	NoteDeliveryNoticeItemSubjectText                                              *string  `json:"NoteDeliveryNoticeItemSubjectText"`
	NoteDeliveryNoticeItemContentText                                              *string  `json:"NoteDeliveryNoticeItemContentText"`
	NoteDeliveryNoticeItemIdentifier                                               *string  `json:"NoteDeliveryNoticeItemIdentifier"`
	ReferencedSalesOrderDocumentIssuerAssignedIdentifier                           *string  `json:"ReferencedSalesOrderDocumentIssuerAssignedIdentifier"`
	ReferencedSalesOrderDocumentItemLineIdentifier                                 *string  `json:"ReferencedSalesOrderDocumentItemLineIdentifier"`
	ReferencedSalesOrderDocumentRevisionIdentifier                                 *string  `json:"ReferencedSalesOrderDocumentRevisionIdentifier"`
	ReferencedOrdersDocumentIssureAssignedIdentifier                               *string  `json:"ReferencedOrdersDocumentIssureAssignedIdentifier"`
	ReferencedOrdersDocumentItemLineIdentifier                                     *string  `json:"ReferencedOrdersDocumentItemLineIdentifier"`
	ReferencedOrdersDocumentRevisionIdentifier                                     *string  `json:"ReferencedOrdersDocumentRevisionIdentifier"`
	TradePriceTypeCode                                                             *string  `json:"TradePriceTypeCode"`
	TradeDeliveryPriceChargeAmount                                                 *float32 `json:"TradeDeliveryPriceChargeAmount"`
	TradePriceBasisQuantity                                                        *float32 `json:"TradePriceBasisQuantity"`
	TradePriceBasisUnitCode                                                        *string  `json:"TradePriceBasisUnitCode"`
	OrderQuantityInBaseUnit                                                        *float32 `json:"OrderQuantityInBaseUnit"`
	DeliveryQuantityInBaseUnit                                                     *float32 `json:"DeliveryQuantityInBaseUnit"`
	SupplyChainTradeDeliveryPackageQuantity                                        *float32 `json:"SupplyChainTradeDeliveryPackageQuantity"`
	SupplyChainTradeDeliveryProductUnitQuantity                                    *float32 `json:"SupplyChainTradeDeliveryProductUnitQuantity"`
	SupplyChainTradeDeliveryPerPackageUnitQuantity                                 *float32 `json:"SupplyChainTradeDeliveryPerPackageUnitQuantity"`
	SupplyChainTradeDeliveryDespatchedQuantity                                     *float32 `json:"SupplyChainTradeDeliveryDespatchedQuantity"`
	SupplyChainTradeDeliveryRequestedQuantity                                      *float32 `json:"SupplyChainTradeDeliveryRequestedQuantity"`
	SupplyChainTradeOrderRequestedPackageQuantity                                  *float32 `json:"SupplyChainTradeOrderRequestedPackageQuantity"`
	SupplyChainTradeDeliveryRequestedPackageQuantity                               *float32 `json:"SupplyChainTradeDeliveryRequestedPackageQuantity"`
	SupplyChainTradeDeliveryRemainingRequestedQuantity                             *float32 `json:"SupplyChainTradeDeliveryRemainingRequestedQuantity"`
	SupplyChainTradeDeliveryBilledQuantity                                         *float32 `json:"SupplyChainTradeDeliveryBilledQuantity"`
	ItemTradeDeliverToPartyIdentifier                                              *string  `json:"ItemTradeDeliverToPartyIdentifier"`
	ItemTradeDeliverToPartyGlobalIdentifier                                        *string  `json:"ItemTradeDeliverToPartyGlobalIdentifier"`
	ItemTradeDeliverToPartyName                                                    *string  `json:"ItemTradeDeliverToPartyName"`
	ItemTradeDeliverToPartyContactPersonName                                       *string  `json:"ItemTradeDeliverToPartyContactPersonName"`
	ItemTradeDeliverToPartyContactDepartmentName                                   *string  `json:"ItemTradeDeliverToPartyContactDepartmentName"`
	ItemDeliverToPartyPhoneNumber                                                  *string  `json:"ItemDeliverToPartyPhoneNumber"`
	ItemDeliverToPartyFaxNumber                                                    *string  `json:"ItemDeliverToPartyFaxNumber"`
	ItemDeliverToPartyEMailAddress                                                 *string  `json:"ItemDeliverToPartyEMailAddress"`
	ItemDeliverToPartyAddressPostalCode                                            *string  `json:"ItemDeliverToPartyAddressPostalCode"`
	ItemDeliverToPartyAddress                                                      *string  `json:"ItemDeliverToPartyAddress"`
	SupplyChainDeliveryEventIdentifier                                             *string  `json:"SupplyChainDeliveryEventIdentifier"`
	SupplyChainDeliveryEventOccurrenceDate                                         *string  `json:"SupplyChainDeliveryEventOccurrenceDate"`
	LastShipToPartyDeliveryDate                                                    *string  `json:"SupplyChainEventTypeCode"`
	SupplyChainEventTypeCode                                                       *string  `json:"LastShipToPartyDeliveryDate"`
	SupplyChainEventRequirementOccurrenceDate                                      *string  `json:"SupplyChainEventRequirementOccurrenceDate"`
	LogisticsLocationIdentification                                                *string  `json:"LogisticsLocationIdentification"`
	LogisticsLocationName                                                          *string  `json:"LogisticsLocationName"`
	SupplyChainEventPlannedOccurrenceDate                                          *string  `json:"SupplyChainEventPlannedOccurrenceDate"`
	TradeTaxTypeCode                                                               *string  `json:"TradeTaxTypeCode"`
	ItemTradeTaxBasisAmount                                                        *float32 `json:"ItemTradeTaxBasisAmount"`
	ItemTradeTaxCategoryCode                                                       *string  `json:"ItemTradeTaxCategoryCode"`
	TradeTaxCategoryName                                                           *string  `json:"TradeTaxCategoryName"`
	ItemTradeTaxRateApplicablePercent                                              *float32 `json:"ItemTradeTaxRateApplicablePercent"`
	ItemTradeTaxGrandTotalAmount                                                   *float32 `json:"ItemTradeTaxGrandTotalAmount"`
	ItemTradeDeliveryNoticeSettlementMonetarySummationNetTotalAmount               *float32 `json:"ItemTradeDeliveryNoticeSettlementMonetarySummationNetTotalAmount"`
	ItemTradeSettlementMonetarySummationTaxAmount                                  *float32 `json:"ItemTradeSettlementMonetarySummationTaxAmount"`
	ItemTradeDeliveryNoticeSettlementMonetarySummationIncludingTaxesNetTotalAmount *float32 `json:"ItemTradeDeliveryNoticeSettlementMonetarySummationIncludingTaxesNetTotalAmount"`
	TradeProductIdentifier                                                         *string  `json:"TradeProductIdentifier"`
	TradeProductGlobalIdentifier                                                   *string  `json:"TradeProductGlobalIdentifier"`
	TradeProductSellerAssignedIdentifier                                           *string  `json:"TradeProductSellerAssignedIdentifier"`
	TradeProductBuyerAssignedIdentifier                                            *string  `json:"TradeProductBuyerAssignedIdentifier"`
	TradeProductManufacturerAssignedIdentifier                                     *string  `json:"TradeProductManufacturerAssignedIdentifier"`
	TradeProductName                                                               *string  `json:"TradeProductName"`
	TradeProductDescription                                                        *string  `json:"TradeProductDescription"`
	TradeProductTypeCode                                                           *string  `json:"TradeProductTypeCode"`
	TradeProductEndItemTypeCode                                                    *string  `json:"TradeProductEndItemTypeCode"`
	TradeProductSizeCode                                                           *string  `json:"TradeProductSizeCode"`
	TradeProductSizeDescriptionText                                                *string  `json:"TradeProductSizeDescriptionText"`
	TradeManufacturerIdentifier                                                    *string  `json:"TradeManufacturerIdentifier"`
	TradeManufacturerName                                                          *string  `json:"TradeManufacturerName"`
	ReferencedLogisticsPackageUnitQuantity                                         *float32 `json:"ReferencedLogisticsPackageUnitQuantity"`
	ReferencedLogisticsPackageQuantityUnitCode                                     *string  `json:"ReferencedLogisticsPackageQuantityUnitCode"`
	ReferencedLogisticsPackageTypeCode                                             *string  `json:"ReferencedLogisticsPackageTypeCode"`
	ReferencedLogisticsPackageIdentifier                                           *string  `json:"ReferencedLogisticsPackageIdentifier"`
}
