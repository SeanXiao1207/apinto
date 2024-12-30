// Code generated by 'yaegi extract github.com/eolinker/eosc'. DO NOT EDIT.

package script_handler

import (
	"go/constant"
	"go/token"
	"os"
	"reflect"

	"github.com/eolinker/eosc"
	"github.com/traefik/yaegi/stdlib"
)

func init() {
	stdlib.Symbols["github.com/eolinker/eosc/eosc"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"ErrorConfigFieldUnknown":              reflect.ValueOf(&eosc.ErrorConfigFieldUnknown).Elem(),
		"ErrorConfigIsNil":                     reflect.ValueOf(&eosc.ErrorConfigIsNil).Elem(),
		"ErrorConfigType":                      reflect.ValueOf(&eosc.ErrorConfigType).Elem(),
		"ErrorDriverNotExist":                  reflect.ValueOf(&eosc.ErrorDriverNotExist).Elem(),
		"ErrorDriverNotMatch":                  reflect.ValueOf(&eosc.ErrorDriverNotMatch).Elem(),
		"ErrorNotAllowCreateForSingleton":      reflect.ValueOf(&eosc.ErrorNotAllowCreateForSingleton).Elem(),
		"ErrorNotGetSillForRequire":            reflect.ValueOf(&eosc.ErrorNotGetSillForRequire).Elem(),
		"ErrorParamNotExist":                   reflect.ValueOf(&eosc.ErrorParamNotExist).Elem(),
		"ErrorParamsIsNil":                     reflect.ValueOf(&eosc.ErrorParamsIsNil).Elem(),
		"ErrorProfessionDependencies":          reflect.ValueOf(&eosc.ErrorProfessionDependencies).Elem(),
		"ErrorProfessionNotExist":              reflect.ValueOf(&eosc.ErrorProfessionNotExist).Elem(),
		"ErrorProfessionNotMatch":              reflect.ValueOf(&eosc.ErrorProfessionNotMatch).Elem(),
		"ErrorRegisterConflict":                reflect.ValueOf(&eosc.ErrorRegisterConflict).Elem(),
		"ErrorRequire":                         reflect.ValueOf(&eosc.ErrorRequire).Elem(),
		"ErrorStoreReadOnly":                   reflect.ValueOf(&eosc.ErrorStoreReadOnly).Elem(),
		"ErrorTargetNotImplementSkill":         reflect.ValueOf(&eosc.ErrorTargetNotImplementSkill).Elem(),
		"ErrorUnsupportedKind":                 reflect.ValueOf(&eosc.ErrorUnsupportedKind).Elem(),
		"ErrorWorkerNotExits":                  reflect.ValueOf(&eosc.ErrorWorkerNotExits).Elem(),
		"ErrorWorkerNotRunning":                reflect.ValueOf(&eosc.ErrorWorkerNotRunning).Elem(),
		"EventDel":                             reflect.ValueOf(constant.MakeFromLiteral("\"delete\"", token.STRING, 0)),
		"EventInit":                            reflect.ValueOf(constant.MakeFromLiteral("\"init\"", token.STRING, 0)),
		"EventReset":                           reflect.ValueOf(constant.MakeFromLiteral("\"getFunc\"", token.STRING, 0)),
		"EventSet":                             reflect.ValueOf(constant.MakeFromLiteral("\"set\"", token.STRING, 0)),
		"File_message_proto":                   reflect.ValueOf(&eosc.File_message_proto).Elem(),
		"NamespaceCustomer":                    reflect.ValueOf(constant.MakeFromLiteral("\"customer\"", token.STRING, 0)),
		"NamespaceExtender":                    reflect.ValueOf(constant.MakeFromLiteral("\"extender\"", token.STRING, 0)),
		"NamespaceProfession":                  reflect.ValueOf(constant.MakeFromLiteral("\"profession\"", token.STRING, 0)),
		"NamespaceVariable":                    reflect.ValueOf(constant.MakeFromLiteral("\"variable\"", token.STRING, 0)),
		"NamespaceWorker":                      reflect.ValueOf(constant.MakeFromLiteral("\"worker\"", token.STRING, 0)),
		"NewExtenderRegister":                  reflect.ValueOf(eosc.NewExtenderRegister),
		"Now":                                  reflect.ValueOf(eosc.Now),
		"ProcessAdmin":                         reflect.ValueOf(constant.MakeFromLiteral("\"admin\"", token.STRING, 0)),
		"ProcessHelper":                        reflect.ValueOf(constant.MakeFromLiteral("\"helper\"", token.STRING, 0)),
		"ProcessMaster":                        reflect.ValueOf(constant.MakeFromLiteral("\"master\"", token.STRING, 0)),
		"ProcessWorker":                        reflect.ValueOf(constant.MakeFromLiteral("\"worker\"", token.STRING, 0)),
		"ProfessionConfig_ProfessionMod_name":  reflect.ValueOf(&eosc.ProfessionConfig_ProfessionMod_name).Elem(),
		"ProfessionConfig_ProfessionMod_value": reflect.ValueOf(&eosc.ProfessionConfig_ProfessionMod_value).Elem(),
		"ProfessionConfig_Singleton":           reflect.ValueOf(eosc.ProfessionConfig_Singleton),
		"ProfessionConfig_Worker":              reflect.ValueOf(eosc.ProfessionConfig_Worker),
		"ReadStringFromEntry":                  reflect.ValueOf(eosc.ReadStringFromEntry),
		"SHA1":                                 reflect.ValueOf(eosc.SHA1),
		"SettingModeBatch":                     reflect.ValueOf(eosc.SettingModeBatch),
		"SettingModeReadonly":                  reflect.ValueOf(eosc.SettingModeReadonly),
		"SettingModeSingleton":                 reflect.ValueOf(eosc.SettingModeSingleton),
		"SplitWorkerId":                        reflect.ValueOf(eosc.SplitWorkerId),
		"String":                               reflect.ValueOf(eosc.String),
		"ToWorkerId":                           reflect.ValueOf(eosc.ToWorkerId),
		"Version":                              reflect.ValueOf(eosc.Version),

		// type definitions
		"DriverConfig":                   reflect.ValueOf((*eosc.DriverConfig)(nil)),
		"EoFiles":                        reflect.ValueOf((*eosc.EoFiles)(nil)),
		"ExtenderBuilder":                reflect.ValueOf((*eosc.ExtenderBuilder)(nil)),
		"ExtenderRegister":               reflect.ValueOf((*eosc.ExtenderRegister)(nil)),
		"ExtendersSettings":              reflect.ValueOf((*eosc.ExtendersSettings)(nil)),
		"FormatterConfig":                reflect.ValueOf((*eosc.FormatterConfig)(nil)),
		"GzipFile":                       reflect.ValueOf((*eosc.GzipFile)(nil)),
		"ICustomerVar":                   reflect.ValueOf((*eosc.ICustomerVar)(nil)),
		"IDataMarshaller":                reflect.ValueOf((*eosc.IDataMarshaller)(nil)),
		"IEntry":                         reflect.ValueOf((*eosc.IEntry)(nil)),
		"IExtenderConfigChecker":         reflect.ValueOf((*eosc.IExtenderConfigChecker)(nil)),
		"IExtenderDriver":                reflect.ValueOf((*eosc.IExtenderDriver)(nil)),
		"IExtenderDriverFactory":         reflect.ValueOf((*eosc.IExtenderDriverFactory)(nil)),
		"IExtenderDriverManager":         reflect.ValueOf((*eosc.IExtenderDriverManager)(nil)),
		"IExtenderDriverRegister":        reflect.ValueOf((*eosc.IExtenderDriverRegister)(nil)),
		"IExtenderDrivers":               reflect.ValueOf((*eosc.IExtenderDrivers)(nil)),
		"IFormatter":                     reflect.ValueOf((*eosc.IFormatter)(nil)),
		"IFormatterFactory":              reflect.ValueOf((*eosc.IFormatterFactory)(nil)),
		"IMetricEntry":                   reflect.ValueOf((*eosc.IMetricEntry)(nil)),
		"IProfession":                    reflect.ValueOf((*eosc.IProfession)(nil)),
		"IProfessions":                   reflect.ValueOf((*eosc.IProfessions)(nil)),
		"IRequires":                      reflect.ValueOf((*eosc.IRequires)(nil)),
		"ISetting":                       reflect.ValueOf((*eosc.ISetting)(nil)),
		"ISettings":                      reflect.ValueOf((*eosc.ISettings)(nil)),
		"IVariable":                      reflect.ValueOf((*eosc.IVariable)(nil)),
		"IWorker":                        reflect.ValueOf((*eosc.IWorker)(nil)),
		"IWorkerDestroy":                 reflect.ValueOf((*eosc.IWorkerDestroy)(nil)),
		"IWorkers":                       reflect.ValueOf((*eosc.IWorkers)(nil)),
		"Item":                           reflect.ValueOf((*eosc.Item)(nil)),
		"ProcessStatus":                  reflect.ValueOf((*eosc.ProcessStatus)(nil)),
		"ProfessionConfig":               reflect.ValueOf((*eosc.ProfessionConfig)(nil)),
		"ProfessionConfig_ProfessionMod": reflect.ValueOf((*eosc.ProfessionConfig_ProfessionMod)(nil)),
		"ProfessionConfigs":              reflect.ValueOf((*eosc.ProfessionConfigs)(nil)),
		"RequireId":                      reflect.ValueOf((*eosc.RequireId)(nil)),
		"SettingMode":                    reflect.ValueOf((*eosc.SettingMode)(nil)),
		"TWorker":                        reflect.ValueOf((*eosc.TWorker)(nil)),
		"WorkerConfig":                   reflect.ValueOf((*eosc.WorkerConfig)(nil)),

		// interface wrapper definitions
		"_ExtenderBuilder":         reflect.ValueOf((*_github_com_eolinker_eosc_ExtenderBuilder)(nil)),
		"_ICustomerVar":            reflect.ValueOf((*_github_com_eolinker_eosc_ICustomerVar)(nil)),
		"_IDataMarshaller":         reflect.ValueOf((*_github_com_eolinker_eosc_IDataMarshaller)(nil)),
		"_IEntry":                  reflect.ValueOf((*_github_com_eolinker_eosc_IEntry)(nil)),
		"_IExtenderConfigChecker":  reflect.ValueOf((*_github_com_eolinker_eosc_IExtenderConfigChecker)(nil)),
		"_IExtenderDriver":         reflect.ValueOf((*_github_com_eolinker_eosc_IExtenderDriver)(nil)),
		"_IExtenderDriverFactory":  reflect.ValueOf((*_github_com_eolinker_eosc_IExtenderDriverFactory)(nil)),
		"_IExtenderDriverManager":  reflect.ValueOf((*_github_com_eolinker_eosc_IExtenderDriverManager)(nil)),
		"_IExtenderDriverRegister": reflect.ValueOf((*_github_com_eolinker_eosc_IExtenderDriverRegister)(nil)),
		"_IExtenderDrivers":        reflect.ValueOf((*_github_com_eolinker_eosc_IExtenderDrivers)(nil)),
		"_IFormatter":              reflect.ValueOf((*_github_com_eolinker_eosc_IFormatter)(nil)),
		"_IFormatterFactory":       reflect.ValueOf((*_github_com_eolinker_eosc_IFormatterFactory)(nil)),
		"_IMetricEntry":            reflect.ValueOf((*_github_com_eolinker_eosc_IMetricEntry)(nil)),
		"_IProfession":             reflect.ValueOf((*_github_com_eolinker_eosc_IProfession)(nil)),
		"_IProfessions":            reflect.ValueOf((*_github_com_eolinker_eosc_IProfessions)(nil)),
		"_IRequires":               reflect.ValueOf((*_github_com_eolinker_eosc_IRequires)(nil)),
		"_ISetting":                reflect.ValueOf((*_github_com_eolinker_eosc_ISetting)(nil)),
		"_ISettings":               reflect.ValueOf((*_github_com_eolinker_eosc_ISettings)(nil)),
		"_IVariable":               reflect.ValueOf((*_github_com_eolinker_eosc_IVariable)(nil)),
		"_IWorker":                 reflect.ValueOf((*_github_com_eolinker_eosc_IWorker)(nil)),
		"_IWorkerDestroy":          reflect.ValueOf((*_github_com_eolinker_eosc_IWorkerDestroy)(nil)),
		"_IWorkers":                reflect.ValueOf((*_github_com_eolinker_eosc_IWorkers)(nil)),
	}
}

// _github_com_eolinker_eosc_ExtenderBuilder is an interface wrapper for ExtenderBuilder type
type _github_com_eolinker_eosc_ExtenderBuilder struct {
	IValue    interface{}
	WRegister func(register eosc.IExtenderDriverRegister)
}

func (W _github_com_eolinker_eosc_ExtenderBuilder) Register(register eosc.IExtenderDriverRegister) {
	W.WRegister(register)
}

// _github_com_eolinker_eosc_ICustomerVar is an interface wrapper for ICustomerVar type
type _github_com_eolinker_eosc_ICustomerVar struct {
	IValue  interface{}
	WExists func(key string, field string) bool
	WGet    func(key string, field string) (string, bool)
	WGetAll func(key string) (map[string]string, bool)
}

func (W _github_com_eolinker_eosc_ICustomerVar) Exists(key string, field string) bool {
	return W.WExists(key, field)
}
func (W _github_com_eolinker_eosc_ICustomerVar) Get(key string, field string) (string, bool) {
	return W.WGet(key, field)
}
func (W _github_com_eolinker_eosc_ICustomerVar) GetAll(key string) (map[string]string, bool) {
	return W.WGetAll(key)
}

// _github_com_eolinker_eosc_IDataMarshaller is an interface wrapper for IDataMarshaller type
type _github_com_eolinker_eosc_IDataMarshaller struct {
	IValue  interface{}
	WEncode func(startIndex int) ([]byte, []*os.File, error)
}

func (W _github_com_eolinker_eosc_IDataMarshaller) Encode(startIndex int) ([]byte, []*os.File, error) {
	return W.WEncode(startIndex)
}

// _github_com_eolinker_eosc_IEntry is an interface wrapper for IEntry type
type _github_com_eolinker_eosc_IEntry struct {
	IValue     interface{}
	WChildren  func(child string) []eosc.IEntry
	WRead      func(pattern string) interface{}
	WReadLabel func(pattern string) string
}

func (W _github_com_eolinker_eosc_IEntry) Children(child string) []eosc.IEntry {
	return W.WChildren(child)
}
func (W _github_com_eolinker_eosc_IEntry) Read(pattern string) interface{} {
	return W.WRead(pattern)
}
func (W _github_com_eolinker_eosc_IEntry) ReadLabel(pattern string) string {
	return W.WReadLabel(pattern)
}

// _github_com_eolinker_eosc_IExtenderConfigChecker is an interface wrapper for IExtenderConfigChecker type
type _github_com_eolinker_eosc_IExtenderConfigChecker struct {
	IValue interface{}
	WCheck func(v interface{}, workers map[eosc.RequireId]eosc.IWorker) error
}

func (W _github_com_eolinker_eosc_IExtenderConfigChecker) Check(v interface{}, workers map[eosc.RequireId]eosc.IWorker) error {
	return W.WCheck(v, workers)
}

// _github_com_eolinker_eosc_IExtenderDriver is an interface wrapper for IExtenderDriver type
type _github_com_eolinker_eosc_IExtenderDriver struct {
	IValue      interface{}
	WConfigType func() reflect.Type
	WCreate     func(id string, name string, v interface{}, workers map[eosc.RequireId]eosc.IWorker) (eosc.IWorker, error)
}

func (W _github_com_eolinker_eosc_IExtenderDriver) ConfigType() reflect.Type {
	return W.WConfigType()
}
func (W _github_com_eolinker_eosc_IExtenderDriver) Create(id string, name string, v interface{}, workers map[eosc.RequireId]eosc.IWorker) (eosc.IWorker, error) {
	return W.WCreate(id, name, v, workers)
}

// _github_com_eolinker_eosc_IExtenderDriverFactory is an interface wrapper for IExtenderDriverFactory type
type _github_com_eolinker_eosc_IExtenderDriverFactory struct {
	IValue  interface{}
	WCreate func(profession string, name string, label string, desc string, params map[string]interface{}) (eosc.IExtenderDriver, error)
	WRender func() interface{}
}

func (W _github_com_eolinker_eosc_IExtenderDriverFactory) Create(profession string, name string, label string, desc string, params map[string]interface{}) (eosc.IExtenderDriver, error) {
	return W.WCreate(profession, name, label, desc, params)
}
func (W _github_com_eolinker_eosc_IExtenderDriverFactory) Render() interface{} {
	return W.WRender()
}

// _github_com_eolinker_eosc_IExtenderDriverManager is an interface wrapper for IExtenderDriverManager type
type _github_com_eolinker_eosc_IExtenderDriverManager struct {
	IValue                  interface{}
	WRegisterExtenderDriver func(name string, factory eosc.IExtenderDriverFactory) error
}

func (W _github_com_eolinker_eosc_IExtenderDriverManager) RegisterExtenderDriver(name string, factory eosc.IExtenderDriverFactory) error {
	return W.WRegisterExtenderDriver(name, factory)
}

// _github_com_eolinker_eosc_IExtenderDriverRegister is an interface wrapper for IExtenderDriverRegister type
type _github_com_eolinker_eosc_IExtenderDriverRegister struct {
	IValue                  interface{}
	WRegisterExtenderDriver func(name string, factory eosc.IExtenderDriverFactory) error
}

func (W _github_com_eolinker_eosc_IExtenderDriverRegister) RegisterExtenderDriver(name string, factory eosc.IExtenderDriverFactory) error {
	return W.WRegisterExtenderDriver(name, factory)
}

// _github_com_eolinker_eosc_IExtenderDrivers is an interface wrapper for IExtenderDrivers type
type _github_com_eolinker_eosc_IExtenderDrivers struct {
	IValue     interface{}
	WGetDriver func(name string) (eosc.IExtenderDriverFactory, bool)
}

func (W _github_com_eolinker_eosc_IExtenderDrivers) GetDriver(name string) (eosc.IExtenderDriverFactory, bool) {
	return W.WGetDriver(name)
}

// _github_com_eolinker_eosc_IFormatter is an interface wrapper for IFormatter type
type _github_com_eolinker_eosc_IFormatter struct {
	IValue  interface{}
	WFormat func(entry eosc.IEntry) []byte
}

func (W _github_com_eolinker_eosc_IFormatter) Format(entry eosc.IEntry) []byte {
	return W.WFormat(entry)
}

// _github_com_eolinker_eosc_IFormatterFactory is an interface wrapper for IFormatterFactory type
type _github_com_eolinker_eosc_IFormatterFactory struct {
	IValue  interface{}
	WCreate func(cfg eosc.FormatterConfig, extendCfg ...interface{}) (eosc.IFormatter, error)
}

func (W _github_com_eolinker_eosc_IFormatterFactory) Create(cfg eosc.FormatterConfig, extendCfg ...interface{}) (eosc.IFormatter, error) {
	return W.WCreate(cfg, extendCfg...)
}

// _github_com_eolinker_eosc_IMetricEntry is an interface wrapper for IMetricEntry type
type _github_com_eolinker_eosc_IMetricEntry struct {
	IValue    interface{}
	WChildren func(child string) []eosc.IMetricEntry
	WGetFloat func(pattern string) (float64, bool)
	WRead     func(pattern string) string
}

func (W _github_com_eolinker_eosc_IMetricEntry) Children(child string) []eosc.IMetricEntry {
	return W.WChildren(child)
}
func (W _github_com_eolinker_eosc_IMetricEntry) GetFloat(pattern string) (float64, bool) {
	return W.WGetFloat(pattern)
}
func (W _github_com_eolinker_eosc_IMetricEntry) Read(pattern string) string {
	return W.WRead(pattern)
}

// _github_com_eolinker_eosc_IProfession is an interface wrapper for IProfession type
type _github_com_eolinker_eosc_IProfession struct {
	IValue      interface{}
	WAppendAttr func() []string
	WDrivers    func() []*eosc.DriverConfig
	WGetDriver  func(name string) (*eosc.DriverConfig, bool)
	WHasDriver  func(name string) bool
	WMod        func() eosc.ProfessionConfig_ProfessionMod
}

func (W _github_com_eolinker_eosc_IProfession) AppendAttr() []string {
	return W.WAppendAttr()
}
func (W _github_com_eolinker_eosc_IProfession) Drivers() []*eosc.DriverConfig {
	return W.WDrivers()
}
func (W _github_com_eolinker_eosc_IProfession) GetDriver(name string) (*eosc.DriverConfig, bool) {
	return W.WGetDriver(name)
}
func (W _github_com_eolinker_eosc_IProfession) HasDriver(name string) bool {
	return W.WHasDriver(name)
}
func (W _github_com_eolinker_eosc_IProfession) Mod() eosc.ProfessionConfig_ProfessionMod {
	return W.WMod()
}

// _github_com_eolinker_eosc_IProfessions is an interface wrapper for IProfessions type
type _github_com_eolinker_eosc_IProfessions struct {
	IValue         interface{}
	WAll           func() []*eosc.ProfessionConfig
	WDelete        func(name string) error
	WGetProfession func(name string) (eosc.IProfession, bool)
	WNames         func() []string
	WReset         func(a0 []*eosc.ProfessionConfig)
	WSet           func(name string, profession *eosc.ProfessionConfig) error
}

func (W _github_com_eolinker_eosc_IProfessions) All() []*eosc.ProfessionConfig {
	return W.WAll()
}
func (W _github_com_eolinker_eosc_IProfessions) Delete(name string) error {
	return W.WDelete(name)
}
func (W _github_com_eolinker_eosc_IProfessions) GetProfession(name string) (eosc.IProfession, bool) {
	return W.WGetProfession(name)
}
func (W _github_com_eolinker_eosc_IProfessions) Names() []string {
	return W.WNames()
}
func (W _github_com_eolinker_eosc_IProfessions) Reset(a0 []*eosc.ProfessionConfig) {
	W.WReset(a0)
}
func (W _github_com_eolinker_eosc_IProfessions) Set(name string, profession *eosc.ProfessionConfig) error {
	return W.WSet(name, profession)
}

// _github_com_eolinker_eosc_IRequires is an interface wrapper for IRequires type
type _github_com_eolinker_eosc_IRequires struct {
	IValue          interface{}
	WDel            func(id string)
	WRequireBy      func(requireId string) []string
	WRequireByCount func(requireId string) int
	WRequires       func(id string) []string
	WSet            func(id string, requires []string)
}

func (W _github_com_eolinker_eosc_IRequires) Del(id string) {
	W.WDel(id)
}
func (W _github_com_eolinker_eosc_IRequires) RequireBy(requireId string) []string {
	return W.WRequireBy(requireId)
}
func (W _github_com_eolinker_eosc_IRequires) RequireByCount(requireId string) int {
	return W.WRequireByCount(requireId)
}
func (W _github_com_eolinker_eosc_IRequires) Requires(id string) []string {
	return W.WRequires(id)
}
func (W _github_com_eolinker_eosc_IRequires) Set(id string, requires []string) {
	W.WSet(id, requires)
}

// _github_com_eolinker_eosc_ISetting is an interface wrapper for ISetting type
type _github_com_eolinker_eosc_ISetting struct {
	IValue      interface{}
	WAllWorkers func() []string
	WCheck      func(cfg interface{}) (profession string, name string, driver string, desc string, err error)
	WConfigType func() reflect.Type
	WGet        func() interface{}
	WMode       func() eosc.SettingMode
	WSet        func(conf interface{}) (err error)
}

func (W _github_com_eolinker_eosc_ISetting) AllWorkers() []string {
	return W.WAllWorkers()
}
func (W _github_com_eolinker_eosc_ISetting) Check(cfg interface{}) (profession string, name string, driver string, desc string, err error) {
	return W.WCheck(cfg)
}
func (W _github_com_eolinker_eosc_ISetting) ConfigType() reflect.Type {
	return W.WConfigType()
}
func (W _github_com_eolinker_eosc_ISetting) Get() interface{} {
	return W.WGet()
}
func (W _github_com_eolinker_eosc_ISetting) Mode() eosc.SettingMode {
	return W.WMode()
}
func (W _github_com_eolinker_eosc_ISetting) Set(conf interface{}) (err error) {
	return W.WSet(conf)
}

// _github_com_eolinker_eosc_ISettings is an interface wrapper for ISettings type
type _github_com_eolinker_eosc_ISettings struct {
	IValue         interface{}
	WCheckVariable func(name string) (err error)
	WGetConfig     func(name string) interface{}
	WGetConfigBody func(name string) ([]byte, bool)
	WGetDriver     func(name string) (eosc.ISetting, bool)
	WSettingWorker func(name string, config []byte) error
	WUpdate        func(name string) (err error)
}

func (W _github_com_eolinker_eosc_ISettings) CheckVariable(name string) (err error) {
	return W.WCheckVariable(name)
}
func (W _github_com_eolinker_eosc_ISettings) GetConfig(name string) interface{} {
	return W.WGetConfig(name)
}
func (W _github_com_eolinker_eosc_ISettings) GetConfigBody(name string) ([]byte, bool) {
	return W.WGetConfigBody(name)
}
func (W _github_com_eolinker_eosc_ISettings) GetDriver(name string) (eosc.ISetting, bool) {
	return W.WGetDriver(name)
}
func (W _github_com_eolinker_eosc_ISettings) SettingWorker(name string, config []byte) error {
	return W.WSettingWorker(name, config)
}
func (W _github_com_eolinker_eosc_ISettings) Update(name string) (err error) {
	return W.WUpdate(name)
}

// _github_com_eolinker_eosc_IVariable is an interface wrapper for IVariable type
type _github_com_eolinker_eosc_IVariable struct {
	IValue          interface{}
	WAll            func() map[string]map[string]string
	WGet            func(id string) (string, bool)
	WGetByNamespace func(namespace string) (map[string]string, bool)
	WLen            func() int
	WRemoveRequire  func(id string)
	WSetByNamespace func(namespace string, variables map[string]string) (affectIds []string, clone map[string]string, err error)
	WSetRequire     func(id string, variables []string)
	WUnmarshal      func(buf []byte, typ reflect.Type) (interface{}, []string, error)
}

func (W _github_com_eolinker_eosc_IVariable) All() map[string]map[string]string {
	return W.WAll()
}
func (W _github_com_eolinker_eosc_IVariable) Get(id string) (string, bool) {
	return W.WGet(id)
}
func (W _github_com_eolinker_eosc_IVariable) GetByNamespace(namespace string) (map[string]string, bool) {
	return W.WGetByNamespace(namespace)
}
func (W _github_com_eolinker_eosc_IVariable) Len() int {
	return W.WLen()
}
func (W _github_com_eolinker_eosc_IVariable) RemoveRequire(id string) {
	W.WRemoveRequire(id)
}
func (W _github_com_eolinker_eosc_IVariable) SetByNamespace(namespace string, variables map[string]string) (affectIds []string, clone map[string]string, err error) {
	return W.WSetByNamespace(namespace, variables)
}
func (W _github_com_eolinker_eosc_IVariable) SetRequire(id string, variables []string) {
	W.WSetRequire(id, variables)
}
func (W _github_com_eolinker_eosc_IVariable) Unmarshal(buf []byte, typ reflect.Type) (interface{}, []string, error) {
	return W.WUnmarshal(buf, typ)
}

// _github_com_eolinker_eosc_IWorker is an interface wrapper for IWorker type
type _github_com_eolinker_eosc_IWorker struct {
	IValue      interface{}
	WCheckSkill func(skill string) bool
	WId         func() string
	WReset      func(conf interface{}, workers map[eosc.RequireId]eosc.IWorker) error
	WStart      func() error
	WStop       func() error
}

func (W _github_com_eolinker_eosc_IWorker) CheckSkill(skill string) bool {
	return W.WCheckSkill(skill)
}
func (W _github_com_eolinker_eosc_IWorker) Id() string {
	return W.WId()
}
func (W _github_com_eolinker_eosc_IWorker) Reset(conf interface{}, workers map[eosc.RequireId]eosc.IWorker) error {
	return W.WReset(conf, workers)
}
func (W _github_com_eolinker_eosc_IWorker) Start() error {
	return W.WStart()
}
func (W _github_com_eolinker_eosc_IWorker) Stop() error {
	return W.WStop()
}

// _github_com_eolinker_eosc_IWorkerDestroy is an interface wrapper for IWorkerDestroy type
type _github_com_eolinker_eosc_IWorkerDestroy struct {
	IValue   interface{}
	WDestroy func() error
}

func (W _github_com_eolinker_eosc_IWorkerDestroy) Destroy() error {
	return W.WDestroy()
}

// _github_com_eolinker_eosc_IWorkers is an interface wrapper for IWorkers type
type _github_com_eolinker_eosc_IWorkers struct {
	IValue interface{}
	WGet   func(id string) (eosc.IWorker, bool)
}

func (W _github_com_eolinker_eosc_IWorkers) Get(id string) (eosc.IWorker, bool) {
	return W.WGet(id)
}
