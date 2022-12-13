package entitys

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/goleak"

	"github.com/yinweli/Mizugo/testdata"
)

func TestEntity(t *testing.T) {
	suite.Run(t, new(SuiteEntity))
}

type SuiteEntity struct {
	suite.Suite
	testdata.TestEnv
	timeout time.Duration
}

func (this *SuiteEntity) SetupSuite() {
	this.Change("test-entitys-entity")
	this.timeout = time.Second
}

func (this *SuiteEntity) TearDownSuite() {
	this.Restore()
}

func (this *SuiteEntity) TearDownTest() {
	goleak.VerifyNone(this.T())
}

func (this *SuiteEntity) TestNewEntity() {
	assert.NotNil(this.T(), NewEntity(EntityID(1)))
}

func (this *SuiteEntity) TestEntity() {
	target := NewEntity(EntityID(1))
	target.initialize()
	assert.Equal(this.T(), EntityID(1), target.EntityID())
	assert.True(this.T(), target.Enable())
	target.finalize()
	assert.False(this.T(), target.Enable())
}

func (this *SuiteEntity) TestModule() {
	target := NewEntity(EntityID(1))
	module1 := newModuleTester(ModuleID(1))
	module2 := newModuleTester(ModuleID(2))
	assert.Nil(this.T(), target.AddModule(module1))
	assert.NotNil(this.T(), target.GetModule(module1.ModuleID()))
	assert.NotNil(this.T(), target.AddModule(module1))
	target.initialize()
	assert.NotNil(this.T(), target.AddModule(module2))
	target.finalize()
}

func (this *SuiteEntity) TestEvent() {
	target := NewEntity(EntityID(1))
	target.initialize()

	eventOnce := "eventOnce"
	paramOnce := "paramOnce"
	validOnce := atomic.Bool{}
	target.SubEvent(eventOnce, func(param any) {
		validOnce.Store(param.(string) == paramOnce)
	})
	target.PubOnceEvent(eventOnce, paramOnce)
	time.Sleep(this.timeout)
	assert.True(this.T(), validOnce.Load())

	eventFixed := "eventFixed"
	paramFixed := "paramFixed"
	validFixed := atomic.Int64{}
	target.SubEvent(eventFixed, func(param any) {
		if param.(string) == paramFixed {
			validFixed.Add(1)
		} // if
	})
	fixed := target.PubFixedEvent(eventFixed, paramFixed, time.Millisecond)
	time.Sleep(this.timeout)
	assert.Greater(this.T(), validFixed.Load(), int64(0))

	fixed.Stop()
	target.finalize()
}

func (this *SuiteEntity) TestInitialize() {
	target := NewEntity(EntityID(1))
	module := newModuleTester(ModuleID(1))
	assert.Nil(this.T(), target.AddModule(module))
	target.initialize()
	time.Sleep(this.timeout)
	target.finalize()
	time.Sleep(this.timeout)

	assert.True(this.T(), module.awake.Load())
	assert.True(this.T(), module.start.Load())
	assert.True(this.T(), module.dispose.Load())
	assert.True(this.T(), module.update.Load())
}
