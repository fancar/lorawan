package band

import (
	"time"

	"github.com/brocaar/lorawan"
)

/*

Частоты инициализации – 866.2, 866.4 RX\TX (аналогично российским 868.9 и 869.1).
Доп. частоты, которые должны спускаться на оконечные устройства после процедуры регистрации:
 865.1, 865.3, 865.5 (в данный момент они не спускаются т.к. являются частью ЧП).
По RX2 всё остаётся:
Rx2: 865.525 - SF9BW125

*/

type kg866CustomBand struct {
	band
}

func (b *kg866CustomBand) Name() string {
	return "KG866CUSTOM"
}

func (b *kg866CustomBand) GetDefaults() Defaults {
	return Defaults{
		RX2Frequency:     865525000,
		RX2DataRate:      3,
		ReceiveDelay1:    time.Second,
		ReceiveDelay2:    time.Second * 2,
		JoinAcceptDelay1: time.Second * 5,
		JoinAcceptDelay2: time.Second * 6,
	}
}

func (b *kg866CustomBand) GetDownlinkTXPower(freq uint32) int {
	return 24
}

func (b *kg866CustomBand) GetDefaultMaxUplinkEIRP() float32 {
	return 16
}

func (b *kg866CustomBand) GetPingSlotFrequency(lorawan.DevAddr, time.Duration) (uint32, error) {
	return 866200000, nil
}

func (b *kg866CustomBand) GetRX1ChannelIndexForUplinkChannelIndex(uplinkChannel int) (int, error) {
	return uplinkChannel, nil
}

func (b *kg866CustomBand) GetRX1FrequencyForUplinkFrequency(uplinkFrequency uint32) (uint32, error) {
	return uplinkFrequency, nil
}

func (b *kg866CustomBand) ImplementsTXParamSetup(protocolVersion string) bool {
	return false
}

func newKg866CustomBand(repeaterCompatible bool) (Band, error) {
	b := kg866CustomBand{
		band: band{
			supportsExtraChannels: true,
			cFListMinDR:           0,
			cFListMaxDR:           5,
			dataRates: map[int]DataRate{
				0: {Modulation: LoRaModulation, SpreadFactor: 12, Bandwidth: 125, uplink: true, downlink: true},
				1: {Modulation: LoRaModulation, SpreadFactor: 11, Bandwidth: 125, uplink: true, downlink: true},
				2: {Modulation: LoRaModulation, SpreadFactor: 10, Bandwidth: 125, uplink: true, downlink: true},
				3: {Modulation: LoRaModulation, SpreadFactor: 9, Bandwidth: 125, uplink: true, downlink: true},
				4: {Modulation: LoRaModulation, SpreadFactor: 8, Bandwidth: 125, uplink: true, downlink: true},
				5: {Modulation: LoRaModulation, SpreadFactor: 7, Bandwidth: 125, uplink: true, downlink: true},
				6: {Modulation: LoRaModulation, SpreadFactor: 7, Bandwidth: 250, uplink: true, downlink: true},
			},
			rx1DataRateTable: map[int][]int{
				0: {0, 0, 0, 0, 0, 0},
				1: {1, 0, 0, 0, 0, 0},
				2: {2, 1, 0, 0, 0, 0},
				3: {3, 2, 1, 0, 0, 0},
				4: {4, 3, 2, 1, 0, 0},
				5: {5, 4, 3, 2, 1, 0},
				6: {6, 5, 4, 3, 2, 1},
			},
			txPowerOffsets: []int{
				0,
				-2,
				-4,
				-6,
				-8,
				-10,
				-12,
				-14,
			},
			uplinkChannels: []Channel{
				{Frequency: 866200000, MinDR: 0, MaxDR: 5, enabled: true},
				{Frequency: 866400000, MinDR: 0, MaxDR: 5, enabled: true},
			},

			downlinkChannels: []Channel{
				{Frequency: 866200000, MinDR: 0, MaxDR: 5, enabled: true},
				{Frequency: 866400000, MinDR: 0, MaxDR: 5, enabled: true},
			},
		},
	}

	if repeaterCompatible {
		b.band.maxPayloadSizePerDR = map[string]map[string]map[int]MaxPayloadSize{
			LoRaWAN_1_0_2: {
				latest: { // LoRaWAN 1.0.2B
					0: {M: 59, N: 51},
					1: {M: 59, N: 51},
					2: {M: 59, N: 51},
					3: {M: 123, N: 115},
					4: {M: 230, N: 222},
					5: {M: 230, N: 222},
				},
			},
			LoRaWAN_1_0_3: {
				latest: { // LoRaWAN 1.0.3A
					0: {M: 59, N: 51},
					1: {M: 59, N: 51},
					2: {M: 59, N: 51},
					3: {M: 123, N: 115},
					4: {M: 230, N: 222},
					5: {M: 230, N: 222},
				},
			},
			LoRaWAN_1_1_0: {
				latest: { // LoRaWAN 1.1.0A, 1.1.0B
					0: {M: 59, N: 51},
					1: {M: 59, N: 51},
					2: {M: 59, N: 51},
					3: {M: 123, N: 115},
					4: {M: 230, N: 222},
					5: {M: 230, N: 222},
				},
			},
			latest: {
				RegParamRevRP002_1_0_0: { // RP002-1.0.0
					0: {M: 59, N: 51},
					1: {M: 59, N: 51},
					2: {M: 59, N: 51},
					3: {M: 123, N: 115},
					4: {M: 230, N: 222},
					5: {M: 230, N: 222},
				},
				latest: { // RP002-1.0.1, RP002-1.0.2, RP002-1.0.3
					0: {M: 59, N: 51},
					1: {M: 59, N: 51},
					2: {M: 59, N: 51},
					3: {M: 123, N: 115},
					4: {M: 230, N: 222},
					5: {M: 230, N: 222},
				},
			},
		}
	} else {
		b.band.maxPayloadSizePerDR = map[string]map[string]map[int]MaxPayloadSize{
			LoRaWAN_1_0_2: {
				latest: { // LoRaWAN 1.0.2B
					0: {M: 59, N: 51},
					1: {M: 59, N: 51},
					2: {M: 59, N: 51},
					3: {M: 123, N: 115},
					4: {M: 250, N: 242},
					5: {M: 250, N: 242},
				},
			},
			LoRaWAN_1_0_3: {
				latest: { // LoRaWAN 1.0.3A
					0: {M: 59, N: 51},
					1: {M: 59, N: 51},
					2: {M: 59, N: 51},
					3: {M: 123, N: 115},
					4: {M: 250, N: 242},
					5: {M: 250, N: 242},
				},
			},
			LoRaWAN_1_1_0: {
				latest: { // LoRaWAN 1.1.0A, 1.1.0B
					0: {M: 59, N: 51},
					1: {M: 59, N: 51},
					2: {M: 59, N: 51},
					3: {M: 123, N: 115},
					4: {M: 250, N: 242},
					5: {M: 250, N: 242},
				},
			},
			latest: {
				RegParamRevRP002_1_0_0: { // RP002-1.0.0
					0: {M: 59, N: 51},
					1: {M: 59, N: 51},
					2: {M: 59, N: 51},
					3: {M: 123, N: 115},
					4: {M: 250, N: 242},
					5: {M: 250, N: 242},
				},
				latest: { // RP002-1.0.1, RP002-1.0.2, RP002-1.0.3
					0: {M: 59, N: 51},
					1: {M: 59, N: 51},
					2: {M: 59, N: 51},
					3: {M: 123, N: 115},
					4: {M: 250, N: 242},
					5: {M: 250, N: 242},
				},
			},
		}
	}

	return &b, nil
}
