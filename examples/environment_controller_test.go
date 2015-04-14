package examples

import (
	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

type EnvironmentControllerFixture struct {
	*gunit.Fixture
	hardware   *FakeHVAC
	controller *EnvironmentController
}

func (self *EnvironmentControllerFixture) Setup() {
	self.hardware = NewFakeHardware()
	self.controller = NewController(self.hardware)
}

func (self *EnvironmentControllerFixture) TestEverythingTurnedOffAtStartup() {
	self.activateAllHardwareComponents()
	self.controller = NewController(self.hardware)
	self.assertAllHardwareComponentsAreDeactivated()
}

func (self *EnvironmentControllerFixture) TestEverythingOffWhenComfortable() {
	self.setupComfortableEnvironment()
	self.assertAllHardwareComponentsAreDeactivated()
}

func (self *EnvironmentControllerFixture) TestCoolerAndBlowerWhenHot() {
	self.setupHotEnvironment()
	self.So(self.hardware.String(), should.Equal, "heater BLOWER COOLER low high")
}

func (self *EnvironmentControllerFixture) TestHeaterAndBlowerWhenCold() {
	self.setupColdEnvironment()
	self.So(self.hardware.String(), should.Equal, "HEATER BLOWER cooler low high")
}

func (self *EnvironmentControllerFixture) setupComfortableEnvironment() {
	self.hardware.SetCurrentTemperature(COMFORTABLE)
	self.controller.Regulate()
}
func (self *EnvironmentControllerFixture) setupHotEnvironment() {
	self.hardware.SetCurrentTemperature(TOO_HOT)
	self.controller.Regulate()
}
func (self *EnvironmentControllerFixture) setupColdEnvironment() {
	self.hardware.SetCurrentTemperature(TOO_COLD)
	self.controller.Regulate()
}

func (self *EnvironmentControllerFixture) activateAllHardwareComponents() {
	self.hardware.ActivateBlower()
	self.hardware.ActivateHeater()
	self.hardware.ActivateCooler()
	self.hardware.ActivateHighTemperatureAlarm()
	self.hardware.ActivateLowTemperatureAlarm()
}

func (self *EnvironmentControllerFixture) assertAllHardwareComponentsAreDeactivated() {
	self.So(self.hardware.String(), should.Equal, "heater blower cooler low high")
}
