package entitys

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/goleak"

	"github.com/yinweli/Mizugo/testdata"
)

func TestModulean(t *testing.T) {
	suite.Run(t, new(SuiteModulean))
}

type SuiteModulean struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteModulean) SetupSuite() {
	this.Change("test-modulean")
}

func (this *SuiteModulean) TearDownSuite() {
	this.Restore()
}

func (this *SuiteModulean) TearDownTest() {
	goleak.VerifyNone(this.T())
}

func (this *SuiteModulean) TestNewModulean() {
	assert.NotNil(this.T(), NewModulean())
}

func (this *SuiteModulean) TestAdd() {
	target := NewModulean()

	assert.Nil(this.T(), target.Add(NewModule(ModuleID(1), "module")))
	assert.NotNil(this.T(), target.Get(ModuleID(1)))

	assert.NotNil(this.T(), target.Add(NewModule(ModuleID(1), "module")))
}

func (this *SuiteModulean) TestDel() {
	target := NewModulean()

	assert.Nil(this.T(), target.Add(NewModule(ModuleID(1), "module")))
	module := target.Del(ModuleID(1))
	assert.NotNil(this.T(), module)
	assert.Equal(this.T(), ModuleID(1), module.ModuleID())
	assert.Equal(this.T(), "module", module.Name())
	assert.Nil(this.T(), target.Get(ModuleID(1)))

	assert.Nil(this.T(), target.Del(ModuleID(1)))
}

func (this *SuiteModulean) TestGet() {
	target := NewModulean()

	assert.Nil(this.T(), target.Add(NewModule(ModuleID(1), "module")))
	module := target.Get(ModuleID(1))
	assert.NotNil(this.T(), module)
	assert.Equal(this.T(), ModuleID(1), module.ModuleID())
	assert.Equal(this.T(), "module", module.Name())

	assert.Nil(this.T(), target.Get(ModuleID(2)))
}

func (this *SuiteModulean) TestAll() {
	target := NewModulean()

	assert.Nil(this.T(), target.Add(NewModule(ModuleID(1), "module1")))
	assert.Nil(this.T(), target.Add(NewModule(ModuleID(2), "module2")))
	module := target.All()
	assert.Len(this.T(), module, 2)
	assert.Equal(this.T(), ModuleID(1), module[0].ModuleID())
	assert.Equal(this.T(), "module1", module[0].Name())
	assert.Equal(this.T(), ModuleID(2), module[1].ModuleID())
	assert.Equal(this.T(), "module2", module[1].Name())
}

func (this *SuiteModulean) TestCount() {
	target := NewModulean()

	assert.Nil(this.T(), target.Add(NewModule(ModuleID(1), "module1")))
	assert.Nil(this.T(), target.Add(NewModule(ModuleID(2), "module2")))
	assert.Equal(this.T(), 2, target.Count())
}
