// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package visualsearch

import original "github.com/Azure/azure-sdk-for-go/services/preview/cognitiveservices/v1.0/visualsearch"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type ImagesClient = original.ImagesClient
type Currency = original.Currency

const (
	AED Currency = original.AED
	AFN Currency = original.AFN
	ALL Currency = original.ALL
	AMD Currency = original.AMD
	ANG Currency = original.ANG
	AOA Currency = original.AOA
	ARS Currency = original.ARS
	AUD Currency = original.AUD
	AWG Currency = original.AWG
	AZN Currency = original.AZN
	BAM Currency = original.BAM
	BBD Currency = original.BBD
	BDT Currency = original.BDT
	BGN Currency = original.BGN
	BHD Currency = original.BHD
	BIF Currency = original.BIF
	BMD Currency = original.BMD
	BND Currency = original.BND
	BOB Currency = original.BOB
	BOV Currency = original.BOV
	BRL Currency = original.BRL
	BSD Currency = original.BSD
	BTN Currency = original.BTN
	BWP Currency = original.BWP
	BYR Currency = original.BYR
	BZD Currency = original.BZD
	CAD Currency = original.CAD
	CDF Currency = original.CDF
	CHE Currency = original.CHE
	CHF Currency = original.CHF
	CHW Currency = original.CHW
	CLF Currency = original.CLF
	CLP Currency = original.CLP
	CNY Currency = original.CNY
	COP Currency = original.COP
	COU Currency = original.COU
	CRC Currency = original.CRC
	CUC Currency = original.CUC
	CUP Currency = original.CUP
	CVE Currency = original.CVE
	CZK Currency = original.CZK
	DJF Currency = original.DJF
	DKK Currency = original.DKK
	DOP Currency = original.DOP
	DZD Currency = original.DZD
	EGP Currency = original.EGP
	ERN Currency = original.ERN
	ETB Currency = original.ETB
	EUR Currency = original.EUR
	FJD Currency = original.FJD
	FKP Currency = original.FKP
	GBP Currency = original.GBP
	GEL Currency = original.GEL
	GHS Currency = original.GHS
	GIP Currency = original.GIP
	GMD Currency = original.GMD
	GNF Currency = original.GNF
	GTQ Currency = original.GTQ
	GYD Currency = original.GYD
	HKD Currency = original.HKD
	HNL Currency = original.HNL
	HRK Currency = original.HRK
	HTG Currency = original.HTG
	HUF Currency = original.HUF
	IDR Currency = original.IDR
	ILS Currency = original.ILS
	INR Currency = original.INR
	IQD Currency = original.IQD
	IRR Currency = original.IRR
	ISK Currency = original.ISK
	JMD Currency = original.JMD
	JOD Currency = original.JOD
	JPY Currency = original.JPY
	KES Currency = original.KES
	KGS Currency = original.KGS
	KHR Currency = original.KHR
	KMF Currency = original.KMF
	KPW Currency = original.KPW
	KRW Currency = original.KRW
	KWD Currency = original.KWD
	KYD Currency = original.KYD
	KZT Currency = original.KZT
	LAK Currency = original.LAK
	LBP Currency = original.LBP
	LKR Currency = original.LKR
	LRD Currency = original.LRD
	LSL Currency = original.LSL
	LYD Currency = original.LYD
	MAD Currency = original.MAD
	MDL Currency = original.MDL
	MGA Currency = original.MGA
	MKD Currency = original.MKD
	MMK Currency = original.MMK
	MNT Currency = original.MNT
	MOP Currency = original.MOP
	MRO Currency = original.MRO
	MUR Currency = original.MUR
	MVR Currency = original.MVR
	MWK Currency = original.MWK
	MXN Currency = original.MXN
	MXV Currency = original.MXV
	MYR Currency = original.MYR
	MZN Currency = original.MZN
	NAD Currency = original.NAD
	NGN Currency = original.NGN
	NIO Currency = original.NIO
	NOK Currency = original.NOK
	NPR Currency = original.NPR
	NZD Currency = original.NZD
	OMR Currency = original.OMR
	PAB Currency = original.PAB
	PEN Currency = original.PEN
	PGK Currency = original.PGK
	PHP Currency = original.PHP
	PKR Currency = original.PKR
	PLN Currency = original.PLN
	PYG Currency = original.PYG
	QAR Currency = original.QAR
	RON Currency = original.RON
	RSD Currency = original.RSD
	RUB Currency = original.RUB
	RWF Currency = original.RWF
	SAR Currency = original.SAR
	SBD Currency = original.SBD
	SCR Currency = original.SCR
	SDG Currency = original.SDG
	SEK Currency = original.SEK
	SGD Currency = original.SGD
	SHP Currency = original.SHP
	SLL Currency = original.SLL
	SOS Currency = original.SOS
	SRD Currency = original.SRD
	SSP Currency = original.SSP
	STD Currency = original.STD
	SYP Currency = original.SYP
	SZL Currency = original.SZL
	THB Currency = original.THB
	TJS Currency = original.TJS
	TMT Currency = original.TMT
	TND Currency = original.TND
	TOP Currency = original.TOP
	TRY Currency = original.TRY
	TTD Currency = original.TTD
	TWD Currency = original.TWD
	TZS Currency = original.TZS
	UAH Currency = original.UAH
	UGX Currency = original.UGX
	USD Currency = original.USD
	UYU Currency = original.UYU
	UZS Currency = original.UZS
	VEF Currency = original.VEF
	VND Currency = original.VND
	VUV Currency = original.VUV
	WST Currency = original.WST
	XAF Currency = original.XAF
	XCD Currency = original.XCD
	XOF Currency = original.XOF
	XPF Currency = original.XPF
	YER Currency = original.YER
	ZAR Currency = original.ZAR
	ZMW Currency = original.ZMW
)

type ErrorCode = original.ErrorCode

const (
	InsufficientAuthorization ErrorCode = original.InsufficientAuthorization
	InvalidAuthorization      ErrorCode = original.InvalidAuthorization
	InvalidRequest            ErrorCode = original.InvalidRequest
	None                      ErrorCode = original.None
	RateLimitExceeded         ErrorCode = original.RateLimitExceeded
	ServerError               ErrorCode = original.ServerError
)

type ErrorSubCode = original.ErrorSubCode

const (
	AuthorizationDisabled   ErrorSubCode = original.AuthorizationDisabled
	AuthorizationExpired    ErrorSubCode = original.AuthorizationExpired
	AuthorizationMissing    ErrorSubCode = original.AuthorizationMissing
	AuthorizationRedundancy ErrorSubCode = original.AuthorizationRedundancy
	Blocked                 ErrorSubCode = original.Blocked
	HTTPNotAllowed          ErrorSubCode = original.HTTPNotAllowed
	NotImplemented          ErrorSubCode = original.NotImplemented
	ParameterInvalidValue   ErrorSubCode = original.ParameterInvalidValue
	ParameterMissing        ErrorSubCode = original.ParameterMissing
	ResourceError           ErrorSubCode = original.ResourceError
	UnexpectedError         ErrorSubCode = original.UnexpectedError
)

type ItemAvailability = original.ItemAvailability

const (
	Discontinued        ItemAvailability = original.Discontinued
	InStock             ItemAvailability = original.InStock
	InStoreOnly         ItemAvailability = original.InStoreOnly
	LimitedAvailability ItemAvailability = original.LimitedAvailability
	OnlineOnly          ItemAvailability = original.OnlineOnly
	OutOfStock          ItemAvailability = original.OutOfStock
	PreOrder            ItemAvailability = original.PreOrder
	SoldOut             ItemAvailability = original.SoldOut
)

type SafeSearch = original.SafeSearch

const (
	Moderate SafeSearch = original.Moderate
	Off      SafeSearch = original.Off
	Strict   SafeSearch = original.Strict
)

type Type = original.Type

const (
	TypeAction                     Type = original.TypeAction
	TypeAggregateOffer             Type = original.TypeAggregateOffer
	TypeCreativeWork               Type = original.TypeCreativeWork
	TypeErrorResponse              Type = original.TypeErrorResponse
	TypeIdentifiable               Type = original.TypeIdentifiable
	TypeImageAction                Type = original.TypeImageAction
	TypeImageEntityAction          Type = original.TypeImageEntityAction
	TypeImageKnowledge             Type = original.TypeImageKnowledge
	TypeImageModuleAction          Type = original.TypeImageModuleAction
	TypeImageObject                Type = original.TypeImageObject
	TypeImageRecipesAction         Type = original.TypeImageRecipesAction
	TypeImageRelatedSearchesAction Type = original.TypeImageRelatedSearchesAction
	TypeImageShoppingSourcesAction Type = original.TypeImageShoppingSourcesAction
	TypeImageTag                   Type = original.TypeImageTag
	TypeIntangible                 Type = original.TypeIntangible
	TypeMediaObject                Type = original.TypeMediaObject
	TypeNormalizedQuadrilateral    Type = original.TypeNormalizedQuadrilateral
	TypeOffer                      Type = original.TypeOffer
	TypeOrganization               Type = original.TypeOrganization
	TypePerson                     Type = original.TypePerson
	TypePoint2D                    Type = original.TypePoint2D
	TypeRecipe                     Type = original.TypeRecipe
	TypeResponse                   Type = original.TypeResponse
	TypeResponseBase               Type = original.TypeResponseBase
	TypeStructuredValue            Type = original.TypeStructuredValue
	TypeThing                      Type = original.TypeThing
)

type TypeBasicPropertiesItem = original.TypeBasicPropertiesItem

const (
	TypeAggregateRating TypeBasicPropertiesItem = original.TypeAggregateRating
	TypePropertiesItem  TypeBasicPropertiesItem = original.TypePropertiesItem
	TypeRating          TypeBasicPropertiesItem = original.TypeRating
)

type BasicAction = original.BasicAction
type Action = original.Action
type AggregateOffer = original.AggregateOffer
type AggregateRating = original.AggregateRating
type BasicCreativeWork = original.BasicCreativeWork
type CreativeWork = original.CreativeWork
type CropArea = original.CropArea
type Error = original.Error
type ErrorResponse = original.ErrorResponse
type Filters = original.Filters
type BasicIdentifiable = original.BasicIdentifiable
type Identifiable = original.Identifiable
type BasicImageAction = original.BasicImageAction
type ImageAction = original.ImageAction
type ImageEntityAction = original.ImageEntityAction
type ImageInfo = original.ImageInfo
type ImageKnowledge = original.ImageKnowledge
type ImageModuleAction = original.ImageModuleAction
type ImageObject = original.ImageObject
type ImageRecipesAction = original.ImageRecipesAction
type ImageRelatedSearchesAction = original.ImageRelatedSearchesAction
type ImageShoppingSourcesAction = original.ImageShoppingSourcesAction
type ImagesImageMetadata = original.ImagesImageMetadata
type ImagesModule = original.ImagesModule
type ImageTag = original.ImageTag
type ImageTagRegion = original.ImageTagRegion
type BasicIntangible = original.BasicIntangible
type Intangible = original.Intangible
type KnowledgeRequest = original.KnowledgeRequest
type BasicMediaObject = original.BasicMediaObject
type MediaObject = original.MediaObject
type NormalizedQuadrilateral = original.NormalizedQuadrilateral
type BasicOffer = original.BasicOffer
type Offer = original.Offer
type Organization = original.Organization
type Person = original.Person
type Point2D = original.Point2D
type BasicPropertiesItem = original.BasicPropertiesItem
type PropertiesItem = original.PropertiesItem
type Query = original.Query
type BasicRating = original.BasicRating
type Rating = original.Rating
type Recipe = original.Recipe
type RecipesModule = original.RecipesModule
type RelatedSearchesModule = original.RelatedSearchesModule
type Request = original.Request
type BasicResponse = original.BasicResponse
type Response = original.Response
type BasicResponseBase = original.BasicResponseBase
type ResponseBase = original.ResponseBase
type BasicStructuredValue = original.BasicStructuredValue
type StructuredValue = original.StructuredValue
type BasicThing = original.BasicThing
type Thing = original.Thing

func New() BaseClient {
	return original.New()
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func NewImagesClient() ImagesClient {
	return original.NewImagesClient()
}
func NewImagesClientWithBaseURI(baseURI string) ImagesClient {
	return original.NewImagesClientWithBaseURI(baseURI)
}
func PossibleCurrencyValues() []Currency {
	return original.PossibleCurrencyValues()
}
func PossibleErrorCodeValues() []ErrorCode {
	return original.PossibleErrorCodeValues()
}
func PossibleErrorSubCodeValues() []ErrorSubCode {
	return original.PossibleErrorSubCodeValues()
}
func PossibleItemAvailabilityValues() []ItemAvailability {
	return original.PossibleItemAvailabilityValues()
}
func PossibleSafeSearchValues() []SafeSearch {
	return original.PossibleSafeSearchValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func PossibleTypeBasicPropertiesItemValues() []TypeBasicPropertiesItem {
	return original.PossibleTypeBasicPropertiesItemValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
