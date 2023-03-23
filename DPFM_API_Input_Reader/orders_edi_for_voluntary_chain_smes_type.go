package dpfm_api_input_reader

type OrdersEDIForVoluntaryChainSMEsSDC struct {
	ConnectionKey       string                                          `json:"connection_key"`
	Result              bool                                            `json:"result"`
	RedisKey            string                                          `json:"redis_key"`
	Filepath            string                                          `json:"filepath"`
	APIStatusCode       int                                             `json:"api_status_code"`
	RuntimeSessionID    string                                          `json:"runtime_session_id"`
	BusinessPartnerID   *int                                            `json:"business_partner"`
	ServiceLabel        string                                          `json:"service_label"`
	APIType             string                                          `json:"api_type"`
	DataConcatenation   OrdersEDIForVoluntaryChainSMEsDataConcatenation `json:"DataConcatenation"`
	APISchema           string                                          `json:"api_schema"`
	Accepter            []string                                        `json:"accepter"`
	Deleted             bool                                            `json:"deleted"`
	APIProcessingResult *bool                                           `json:"api_processing_result"`
	APIProcessingError  string                                          `json:"api_processing_error"`
}

type OrdersEDIForVoluntaryChainSMEsDataConcatenation struct {
	Header OrdersEDIForVoluntaryChainSMEsHeader `json:"OrdersHeader"`
	Item   []OrdersEDIForVoluntaryChainSMEsItem `json:"OrdersItem"`
}

type OrdersEDIForVoluntaryChainSMEsHeader struct {
	ExchangedOrdersDocumentIdentifier                        string   `json:"ExchangedOrdersDocumentIdentifier"`
	OrdersDocument                                           *string  `json:"OrdersDocument"`
	ExchangedDocumentContextSpecifiedTransactionIdentifier   *string  `json:"ExchangedDocumentContextSpecifiedTransactionIdentifier"`
	ExchangedOrdersDocumentName                              *string  `json:"ExchangedOrdersDocumentName"`
	ExchangedOrdersDocumentTypeCode                          *string  `json:"ExchangedOrdersDocumentTypeCode"`
	ExchangedOrdersDocumentIssueDate                         *string  `json:"ExchangedOrdersDocumentIssueDate"`
	ExchangedOrdersDocumentPurposeCode                       *string  `json:"ExchangedOrdersDocumentPurposeCode"`
	ExchangedOrdersDocumentRevisionDate                      *string  `json:"ExchangedOrdersDocumentRevisionDate"`
	ExchangedOrdersDocumentRevisionIdentifier                *string  `json:"ExchangedOrdersDocumentRevisionIdentifier"`
	ExchangedOrdersDocumentStatusCode                        *string  `json:"ExchangedOrdersDocumentStatusCode"`
	ExchangedOrdersDocumentSubtypeCode                       *string  `json:"ExchangedOrdersDocumentSubtypeCode"`
	NoteOrdersSubjectText                                    *string  `json:"NoteOrdersSubjectText"`
	NoteOrdersContentText                                    *string  `json:"NoteOrdersContentText"`
	NoteOrdersIdentifier                                     *string  `json:"NoteOrdersIdentifier"`
	SpecifiedBinaryFileIdentifier                            *string  `json:"SpecifiedBinaryFileIdentifier"`
	SpecifiedBinaryFileVersionIdentifier                     *string  `json:"SpecifiedBinaryFileVersionIdentifier"`
	SpecifiedBinaryFileNameText                              *string  `json:"SpecifiedBinaryFileNameText"`
	SpecifiedBineryFileURIIdentifier                         *string  `json:"SpecifiedBineryFileURIIdentifier"`
	SpecifiedBineryFileMIMECode                              *string  `json:"SpecifiedBineryFileMIMECode"`
	SpecifiedBineryFileEncodingCode                          *string  `json:"SpecifiedBineryFileEncodingCode"`
	SpecifiedBineryFileCode                                  *string  `json:"SpecifiedBineryFileCode"`
	SpecifiedBinaryFileDescriptionText                       *string  `json:"SpecifiedBinaryFileDescriptionText"`
	TradeSellerIdentifier                                    *string  `json:"TradeSellerIdentifier"`
	TradeSellerGlobalIdentifier                              *string  `json:"TradeSellerGlobalIdentifier"`
	TradeSellerName                                          *string  `json:"TradeSellerName"`
	TradeBillFromPartyRegisteredIdentifier                   *string  `json:"TradeBillFromPartyRegisteredIdentifier"`
	TradeContactSellerIdentifier                             *string  `json:"TradeContactSellerIdentifier"`
	TradeContactSellerPersonName                             *string  `json:"TradeContactSellerPersonName"`
	TradeContactSellerDepartmentName                         *string  `json:"TradeContactSellerDepartmentName"`
	SellerTelephoneNumber                                    *string  `json:"SellerTelephoneNumber"`
	SellerFaxNumber                                          *string  `json:"SellerFaxNumber"`
	SellerEmailAddress                                       *string  `json:"SellerEmailAddress"`
	SellerAddressPostalCode                                  *string  `json:"SellerAddressPostalCode"`
	SellerAddress                                            *string  `json:"SellerAddress"`
	TradeBuyerIdentifier                                     *string  `json:"TradeBuyerIdentifier"`
	TradeBuyerGlobalIdentifier                               *string  `json:"TradeBuyerGlobalIdentifier"`
	TradeBuyerName                                           *string  `json:"TradeBuyerName"`
	TradeContactBuyerIdentifier                              *string  `json:"TradeContactBuyerIdentifier"`
	TradeContactBuyerPersonName                              *string  `json:"TradeContactBuyerPersonName"`
	TradeContactBuyerDepartmentName                          *string  `json:"TradeContactBuyerDepartmentName"`
	BuyerTelephoneNumber                                     *string  `json:"BuyerTelephoneNumber"`
	BuyerFaxNumber                                           *string  `json:"BuyerFaxNumber"`
	BuyerEmailAddress                                        *string  `json:"BuyerEmailAddress"`
	BuyerAddressPostalCode                                   *string  `json:"BuyerAddressPostalCode"`
	BuyerAddress                                             *string  `json:"BuyerAddress"`
	ReferencedQuotationReplyDocumentIssureAssignedIdentifier *string  `json:"ReferencedQuotationReplyDocumentIssureAssignedIdentifier"`
	ReferencedQuotationReplyDocumentIssueDate                *string  `json:"ReferencedQuotationReplyDocumentIssueDate"`
	ReferencedQuotationReplyDocumentRevisionIdentifier       *string  `json:"ReferencedQuotationReplyDocumentRevisionIdentifier"`
	ReferencedQuotationReplyDocumentInformationText          *string  `json:"ReferencedQuotationReplyDocumentInformationText"`
	ReferencedQuotationReplyDocumentPurposeCode              *string  `json:"ReferencedQuotationReplyDocumentPurposeCode"`
	ReferencedQuotationReplyDocumentSubtypeCode              *string  `json:"ReferencedQuotationReplyDocumentSubtypeCode"`
	ReferencedSalesOrderDocumentIssureAddignedIdentifier     *string  `json:"ReferencedSalesOrderDocumentIssureAddignedIdentifier"`
	ReferencedSalesOrderDocumentIssueDate                    *string  `json:"ReferencedSalesOrderDocumentIssueDate"`
	ReferencedSalesOrderDocumentRevisionIdentifier           *string  `json:"ReferencedSalesOrderDocumentRevisionIdentifier"`
	ReferencedSalesOrderDocumentInformationText              *string  `json:"ReferencedSalesOrderDocumentInformationText"`
	ReferencedSalesOrderDocumentPurposeCode                  *string  `json:"ReferencedSalesOrderDocumentPurposeCode"`
	ReferencedSalesOrderDocumentSubtypeCode                  *string  `json:"ReferencedSalesOrderDocumentSubtypeCode"`
	RelevantTradePartyIdentifier                             *string  `json:"RelevantTradePartyIdentifier"`
	RelevantTradePartyGlobalIdentifier                       *string  `json:"RelevantTradePartyGlobalIdentifier"`
	RelevantTradePartyName                                   *string  `json:"RelevantTradePartyName"`
	RelevantTradePartyRoleCode                               *string  `json:"RelevantTradePartyRoleCode"`
	RelevantTradeContactCode                                 *string  `json:"RelevantTradeContactCode"`
	RelevantTradeContactPersonName                           *string  `json:"RelevantTradeContactPersonName"`
	ProjectIdentifier                                        *string  `json:"ProjectIdentifier"`
	ProjectName                                              *string  `json:"ProjectName"`
	InspectionEventTypeCode                                  *string  `json:"InspectionEventTypeCode"`
	InspectionEventDescriptionText                           *string  `json:"InspectionEventDescriptionText"`
	ProjectPeriodStartDate                                   *string  `json:"ProjectPeriodStartDate"`
	ProjectPeriodEndDate                                     *string  `json:"ProjectPeriodEndDate"`
	TradeShipToPartyIdentifier                               *string  `json:"TradeShipToPartyIdentifier"`
	TradeShipToPartyGlobalIdentifier                         *string  `json:"TradeShipToPartyGlobalIdentifier"`
	TradeShipToPartyName                                     *string  `json:"TradeShipToPartyName"`
	TradeShipToPartyContactIdentifier                        *string  `json:"TradeShipToPartyContactIdentifier"`
	TradeShipToPartyContactPersonName                        *string  `json:"TradeShipToPartyContactPersonName"`
	TradeShipToPartyContactDepartmentName                    *string  `json:"TradeShipToPartyContactDepartmentName"`
	TradeShipToPartyContactPersonIdentifier                  *string  `json:"TradeShipToPartyContactPersonIdentifier"`
	ShipToPartyTelephoneNumber                               *string  `json:"ShipToPartyTelephoneNumber"`
	ShipToPartyFaxNumber                                     *string  `json:"ShipToPartyFaxNumber"`
	ShipToPartyEmailAddress                                  *string  `json:"ShipToPartyEmailAddress"`
	ShipToPartyAddressPostalCode                             *string  `json:"ShipToPartyAddressPostalCode"`
	ShipToPartyAddress                                       *string  `json:"ShipToPartyAddress"`
	LastShipToPartyIdentifier                                *string  `json:"LastShipToPartyIdentifier"`
	TradeShipFromPartyIdentifier                             *string  `json:"TradeShipFromPartyIdentifier"`
	TradeShipFromPartyName                                   *string  `json:"TradeShipFromPartyName"`
	SupplyChainOperationEventIdentifier                      *string  `json:"SupplyChainOperationEventIdentifier"`
	SupplyChainOperationEventOccurrenceDate                  *string  `json:"SupplyChainOperationEventOccurrenceDate"`
	SupplyChainEventDeliveryTypeCode                         *string  `json:"SupplyChainEventDeliveryTypeCode"`
	SupplyChainEventDeliveryDescription                      *string  `json:"SupplyChainEventDeliveryDescription"`
	InstructedTemperatureControlCode                         *string  `json:"InstructedTemperatureControlCode"`
	SupplyChainTradeCurrencyCode                             *string  `json:"SupplyChainTradeCurrencyCode"`
	TradeTaxCalculatedAmount                                 *float32 `json:"TradeTaxCalculatedAmount"`
	TradeTaxTypeCode                                         *string  `json:"TradeTaxTypeCode"`
	TradeTaxBasisAmount                                      *float32 `json:"TradeTaxBasisAmount"`
	TradeTaxCategoryCode                                     *string  `json:"TradeTaxCategoryCode"`
	TradeTaxCategoryName                                     *string  `json:"TradeTaxCategoryName"`
	TradeTaxRatePercent                                      *string  `json:"TradeTaxRatePercent"`
	TradeTaxGrandTotalAmount                                 *float32 `json:"TradeTaxGrandTotalAmount"`
	TradeTaxCalculationMethod                                *string  `json:"TradeTaxCalculationMethod"`
	TradePaymentTermsDescriptionText                         *string  `json:"TradePaymentTermsDescriptionText"`
	TradePaymentTermsTypeCode                                *string  `json:"TradePaymentTermsTypeCode"`
	TradeSettlementMonetarySummationTotalTaxAmount           *float32 `json:"TradeSettlementMonetarySummationTotalTaxAmount"`
	TradeOrdersSettlementMonetarySummationGrandTotalAmount   *float32 `json:"TradeOrdersSettlementMonetarySummationGrandTotalAmount"`
	TradeOrdersSettlementMonetarySummationNetTotalAmount     *float32 `json:"TradeOrdersSettlementMonetarySummationNetTotalAmount"`
	TradeOrdersMonetarySummationIncludingTaxesTotalAmount    *float32 `json:"TradeOrdersMonetarySummationIncludingTaxesTotalAmount"`
}

type OrdersEDIForVoluntaryChainSMEsItem struct {
	ExchangedOrdersDocumentIdentifier                                      string   `json:"ExchangedOrdersDocumentIdentifier"`
	OrdersDocumentItemlineIdentifier                                       string   `json:"OrdersDocumentItemlineIdentifier"`
	OrdersDocumentItemlineStatusCode                                       *string  `json:"OrdersDocumentItemlineStatusCode"`
	OrdersDocumentItemlineStatusReasonCode                                 *string  `json:"OrdersDocumentItemlineStatusReasonCode"`
	OrdersDocumentItemIdentifier                                           *string  `json:"OrdersDocumentItemIdentifier"`
	OrdersDocumentItemBuyerAssignedCategoryCode                            *string  `json:"OrdersDocumentItemBuyerAssignedCategoryCode"`
	NoteOrdersItemItemSubjectText                                          *string  `json:"NoteOrdersItemItemSubjectText"`
	NoteOrdersItemContentText                                              *string  `json:"NoteOrdersItemContentText"`
	NoteOrdersItemIdentifier                                               *string  `json:"NoteOrdersItemIdentifier"`
	ReferencedQuotationReplyDocumentItemIssuerAssignedIdentifier           *string  `json:"ReferencedQuotationReplyDocumentItemIssuerAssignedIdentifier"`
	ReferencedQuotationReplyDocumentItemLineIdentifier                     *string  `json:"ReferencedQuotationReplyDocumentItemLineIdentifier"`
	ReferencedQuotationReplyDocumentItemRevisionIdentifier                 *string  `json:"ReferencedQuotationReplyDocumentItemRevisionIdentifier"`
	ReferencedQuotationReplyDocumentItemInformationText                    *string  `json:"ReferencedQuotationReplyDocumentItemInformationText"`
	ReferencedDocumentItemIssureAssignedIdentifier                         *string  `json:"ReferencedDocumentItemIssureAssignedIdentifier"`
	TradePriceTypeCode                                                     *string  `json:"TradePriceTypeCode"`
	TradeOrdersPriceChargeAmount                                           *float32 `json:"TradeOrdersPriceChargeAmount"`
	TradePriceBasisQuantity                                                *float32 `json:"TradePriceBasisQuantity"`
	TradePriceBasisUnitCode                                                *string  `json:"TradePriceBasisUnitCode"`
	TradeDeliveryTermsDescriptionText                                      *string  `json:"TradeDeliveryTermsDescriptionText"`
	OrderQuantityInBaseUnit                                                *float32 `json:"OrderQuantityInBaseUnit"`
	SupplyChainTradeDeliveryRequestedQuantity                              *float32 `json:"SupplyChainTradeDeliveryRequestedQuantity"`
	SupplyChainTradeDeliveryPerPackageUnitQuantity                         *float32 `json:"SupplyChainTradeDeliveryPerPackageUnitQuantity"`
	SupplyChainTradeOrderRequestedPackageQuantity                          *float32 `json:"SupplyChainTradeOrderRequestedPackageQuantity"`
	SupplyChainTradeDeliveryPackageQuantity                                *float32 `json:"SupplyChainTradeDeliveryPackageQuantity"`
	SupplyChainTradeDeliveryProductUnitQuantity                            *float32 `json:"SupplyChainTradeDeliveryProductUnitQuantity"`
	QuantityUnitCode                                                       *string  `json:"QuantityUnitCode"`
	ItemTradeDeliverToPartyIdentifier                                      *string  `json:"ItemTradeDeliverToPartyIdentifier"`
	ItemTradeDeliverToPartyGlobalIdentifier                                *string  `json:"ItemTradeDeliverToPartyGlobalIdentifier"`
	ItemTradeDeliverToPartyName                                            *string  `json:"ItemTradeDeliverToPartyName"`
	ItemTradeDeliverToPartyContactPersonName                               *string  `json:"ItemTradeDeliverToPartyContactPersonName"`
	ItemTradeDeliverToPartyContactDepartment                               *string  `json:"ItemTradeDeliverToPartyContactDepartment"`
	ItemDeliverToPartyPhoneNumber                                          *string  `json:"ItemDeliverToPartyPhoneNumber"`
	ItemDeliverToPartyFaxNumber                                            *string  `json:"ItemDeliverToPartyFaxNumber"`
	ItemDeliverToPartyEMailAddress                                         *string  `json:"ItemDeliverToPartyEMailAddress"`
	ItemDeliverToPartyAddressPostalCode                                    *string  `json:"ItemDeliverToPartyAddressPostalCode"`
	ItemDeliverToPartyAddress                                              *string  `json:"ItemDeliverToPartyAddress"`
	LastShipToPartyDeliveryDate                                            *string  `json:"LastShipToPartyDeliveryDate"`
	DeliverInstructionsHandlingCode                                        *string  `json:"DeliverInstructionsHandlingCode"`
	SupplyChainEventRequirementOccurrenceDate                              *string  `json:"SupplyChainEventRequirementOccurrenceDate"`
	SupplyChainEventTypeCode                                               *string  `json:"SupplyChainEventTypeCode"`
	SupplyChainEventRequirementOccurrenceTime                              *string  `json:"SupplyChainEventRequirementOccurrenceTime"`
	LogisticsLocationIdentification                                        *string  `json:"LogisticsLocationIdentification"`
	LogisticsLocationName                                                  *string  `json:"LogisticsLocationName"`
	TradeTaxTypeCode                                                       *string  `json:"TradeTaxTypeCode"`
	ItemTradeTaxBasisAmount                                                *float32 `json:"ItemTradeTaxBasisAmount"`
	ItemTradeTaxCategoryCode                                               *string  `json:"ItemTradeTaxCategoryCode"`
	TradeTaxCategoryName                                                   *string  `json:"TradeTaxCategoryName"`
	ItemTradeTaxRateApplicablePercent                                      *float32 `json:"ItemTradeTaxRateApplicablePercent"`
	ItemTradeTaxGrandTotalAmount                                           *float32 `json:"ItemTradeTaxGrandTotalAmount"`
	ItemTradeOrdersSettlementMonetarySummationNetTotalAmount               *float32 `json:"ItemTradeOrdersSettlementMonetarySummationNetTotalAmount"`
	ItemTradeSettlementMonetarySummationTaxAmount                          *float32 `json:"ItemTradeSettlementMonetarySummationTaxAmount"`
	ItemTradeOrdersSettlementMonetarySummationIncludingTaxesNetTotalAmount *float32 `json:"ItemTradeOrdersSettlementMonetarySummationIncludingTaxesNetTotalAmount"`
	TradeProductIdentifier                                                 *string  `json:"TradeProductIdentifier"`
	TradeProductGlobalIdentifier                                           *string  `json:"TradeProductGlobalIdentifier"`
	TradeProductSellerAssignedIdentifier                                   *string  `json:"TradeProductSellerAssignedIdentifier"`
	TradeProductBuyerAssignedIdentifier                                    *string  `json:"TradeProductBuyerAssignedIdentifier"`
	TradeProductManufacturerAssignedIdentifier                             *string  `json:"TradeProductManufacturerAssignedIdentifier"`
	TradeProductName                                                       *string  `json:"TradeProductName"`
	TradeProductDescription                                                *string  `json:"TradeProductDescription"`
	TradeProductTypeCode                                                   *string  `json:"TradeProductTypeCode"`
	TradeProductEndItemTypeCode                                            *string  `json:"TradeProductEndItemTypeCode"`
	TradeProductSizeCode                                                   *string  `json:"TradeProductSizeCode"`
	TradeProductSizeDescriptionText                                        *string  `json:"TradeProductSizeDescriptionText"`
	ProductCharacteristicIdentifier                                        *string  `json:"ProductCharacteristicIdentifier"`
	ProductCharacteristicTypeCode                                          *string  `json:"ProductCharacteristicTypeCode"`
	ProductCharacteristicDescriptionText                                   *string  `json:"ProductCharacteristicDescriptionText"`
	ProductCharacteristicContentTypeCode                                   *string  `json:"ProductCharacteristicContentTypeCode"`
	TradeProductInstanceGlobalSerialIdentifier                             *string  `json:"TradeProductInstanceGlobalSerialIdentifier"`
	TradeProductInstanceBatchIdentifier                                    *string  `json:"TradeProductInstanceBatchIdentifier"`
	ReferenceQualityInspectionDocumentInformationText                      *string  `json:"ReferenceQualityInspectionDocumentInformationText"`
	TradeManufacturerIdentifier                                            *string  `json:"TradeManufacturerIdentifier"`
	TradeManufacturerName                                                  *string  `json:"TradeManufacturerName"`
	ReferencedTechnicalDocumentIssuerAssignedIdentifier                    *string  `json:"ReferencedTechnicalDocumentIssuerAssignedIdentifier"`
	ReferencedTechnicalDocumentRevisionIdentifier                          *string  `json:"ReferencedTechnicalDocumentRevisionIdentifier"`
	ReferencedTechnicalDocumentName                                        *string  `json:"ReferencedTechnicalDocumentName"`
	ReferencedTechnicalDocumentInformationText                             *string  `json:"ReferencedTechnicalDocumentInformationText"`
	ReferencedTechnicalDocumentAttachment                                  *string  `json:"ReferencedTechnicalDocumentAttachment"`
	ReferencedSupplyDocumentIssuerAssignedIdentifier                       *string  `json:"ReferencedSupplyDocumentIssuerAssignedIdentifier"`
	ReferencedSupplyDocumentAttachmentBinaryObject                         *string  `json:"ReferencedSupplyDocumentAttachmentBinaryObject"`
	ReferencedLogisticsPackageUnitQuantity                                 *float32 `json:"ReferencedLogisticsPackageUnitQuantity"`
	ReferencedLogisticsPackageQuantityUnitCode                             *string  `json:"ReferencedLogisticsPackageQuantityUnitCode"`
	ReferencedLogisticsPackageTypeCode                                     *string  `json:"ReferencedLogisticsPackageTypeCode"`
	GoodsProductionIdentifier                                              *string  `json:"GoodsProductionIdentifier"`
	GoodsProductionManufacturingProcessDescriptionText                     *string  `json:"GoodsProductionManufacturingProcessDescriptionText"`
}