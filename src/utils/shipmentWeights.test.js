import {
  calculateNetWeightForProGearWeightTicket,
  calculateNetWeightForWeightTicket,
  calculateNonPPMShipmentNetWeight,
  calculatePPMShipmentNetWeight,
  calculateShipmentNetWeight,
  calculateTotalNetWeightForProGearWeightTickets,
  calculateTotalNetWeightForWeightTickets,
  shipmentIsOverweight,
} from './shipmentWeights';
import { createCompleteProGearWeightTicket } from './test/factories/proGearWeightTicket';
import { createCompleteWeightTicket } from './test/factories/weightTicket';

describe('shipmentWeights utils', () => {
  describe('shipmentIsOverweight', () => {
    it('returns true when the shipment weight is over 110% of the estimated weight', () => {
      expect(shipmentIsOverweight(100, 111)).toEqual(true);
    });

    it('returns false when shipment weight is less than  110% of the estimated weight', () => {
      expect(shipmentIsOverweight(100, 101)).toEqual(false);
    });

    it('returns false when estimated weight is undefined', () => {
      expect(shipmentIsOverweight(undefined, 100)).toEqual(false);
    });
  });
});

describe('calculateNetWeightForWeightTicket', () => {
  it.each([
    [0, 400, 400],
    [15000, 18000, 3000],
    [null, 1500, 0],
    [0, null, 0],
    [null, null, 0],
    [undefined, 1500, 0],
    [0, undefined, 0],
    [undefined, undefined, 0],
    ['not a number', 1500, 0],
    [0, 'not a number', 0],
    ['not a number', 'not a number', 0],
  ])(
    `calculates net weight properly | emptyWeight: %s | fullWeight: %s | expectedNetWeight: %s`,
    (emptyWeight, fullWeight, expectedNetWeight) => {
      const weightTicket = createCompleteWeightTicket(
        {},
        {
          emptyWeight,
          fullWeight,
        },
      );

      expect(calculateNetWeightForWeightTicket(weightTicket)).toEqual(expectedNetWeight);
    },
  );
});

describe('calculateTotalNetWeightForWeightTickets', () => {
  it.each([
    [[{ emptyWeight: 0, fullWeight: 400 }], 400],
    [
      [
        { emptyWeight: 0, fullWeight: 400 },
        { emptyWeight: 15000, fullWeight: 18000 },
      ],
      3400,
    ],
    [
      [
        { emptyWeight: null, fullWeight: 400 },
        { emptyWeight: 14000, fullWeight: 17000 },
      ],
      3000,
    ],
    [
      [
        { emptyWeight: 0, fullWeight: null },
        { emptyWeight: 14000, fullWeight: 19000 },
      ],
      5000,
    ],
    [
      [
        { emptyWeight: null, fullWeight: null },
        { emptyWeight: 14000, fullWeight: 18000 },
      ],
      4000,
    ],
    [
      [
        { emptyWeight: null, fullWeight: null },
        { emptyWeight: null, fullWeight: null },
      ],
      0,
    ],
    [
      [
        { emptyWeight: undefined, fullWeight: 400 },
        { emptyWeight: 14000, fullWeight: 17000 },
      ],
      3000,
    ],
    [
      [
        { emptyWeight: 0, fullWeight: undefined },
        { emptyWeight: 14000, fullWeight: 19000 },
      ],
      5000,
    ],
    [
      [
        { emptyWeight: undefined, fullWeight: undefined },
        { emptyWeight: 14000, fullWeight: 18000 },
      ],
      4000,
    ],
    [
      [
        { emptyWeight: undefined, fullWeight: undefined },
        { emptyWeight: undefined, fullWeight: undefined },
      ],
      0,
    ],
    [
      [
        { emptyWeight: 'not a number', fullWeight: 400 },
        { emptyWeight: 14000, fullWeight: 17000 },
      ],
      3000,
    ],
    [
      [
        { emptyWeight: 0, fullWeight: 'not a number' },
        { emptyWeight: 14000, fullWeight: 19000 },
      ],
      5000,
    ],
    [
      [
        { emptyWeight: 'not a number', fullWeight: 'not a number' },
        { emptyWeight: 14000, fullWeight: 18000 },
      ],
      4000,
    ],
    [
      [
        { emptyWeight: 'not a number', fullWeight: 'not a number' },
        { emptyWeight: 'not a number', fullWeight: 'not a number' },
      ],
      0,
    ],
    [[], 0],
  ])(`calculates total net weight properly`, (weightTicketsFields, expectedNetWeight) => {
    const weightTickets = [];

    weightTicketsFields.forEach((fieldOverrides) => {
      weightTickets.push(createCompleteWeightTicket({}, fieldOverrides));
    });

    expect(calculateTotalNetWeightForWeightTickets(weightTickets)).toEqual(expectedNetWeight);
  });
});

describe('calculateNetWeightForProGearWeightTicket', () => {
  it.each([
    [0, 0],
    [15000, 15000],
    [null, 0],
    [undefined, 0],
    ['not a number', 0],
  ])(
    `calculates net weight properly | emptyWeight: %s | fullWeight: %s | constructedWeight: %s | expectedNetWeight: %s`,
    (weight, expectedNetWeight) => {
      const proGearWeightTicket = createCompleteProGearWeightTicket(
        {},
        {
          weight,
        },
      );

      expect(calculateNetWeightForProGearWeightTicket(proGearWeightTicket)).toEqual(expectedNetWeight);
    },
  );
});

describe('calculateTotalNetWeightForProGearWeightTickets', () => {
  it.each([
    [[{ weight: 0 }], 0],
    [[{ weight: 0 }, { weight: 15000 }], 15000],
    [[{ weight: null }], 0],
    [[{ weight: null }, { weight: 15000 }], 15000],
    [[{ weight: undefined }], 0],
    [[{ weight: undefined }, { weight: 15000 }], 15000],
    [[{ weight: 'not a number' }], 0],
    [[{ weight: 'not a number' }, { weight: 15000 }], 15000],
    [[], 0],
  ])(`calculates total net weight properly`, (proGearWeightTicketsFields, expectedNetWeight) => {
    const proGearWeightTickets = [];

    proGearWeightTicketsFields.forEach((fieldOverrides) => {
      proGearWeightTickets.push(createCompleteProGearWeightTicket({}, fieldOverrides));
    });

    expect(calculateTotalNetWeightForProGearWeightTickets(proGearWeightTickets)).toEqual(expectedNetWeight);
  });
});

describe('Calculating shipment weights', () => {
  const ppmShipments = [
    {
      ppmShipment: {
        weightTickets: [{ emptyWeight: 14000, fullWeight: 19000 }],
      },
    },
    {
      ppmShipment: {
        weightTickets: [
          { emptyWeight: 14000, fullWeight: 19000 },
          { emptyWeight: 12000, fullWeight: 18000 },
        ],
      },
    },
    {
      ppmShipment: {
        weightTickets: [
          { emptyWeight: 14000, fullWeight: 19000 },
          { emptyWeight: 12000, fullWeight: 18000 },
          { emptyWeight: 10000, fullWeight: 20000 },
        ],
      },
    },
  ];

  const hhgShipments = [
    {
      primeActualWeight: 10,
      reweigh: {
        weight: 5,
      },
    },
    {
      primeActualWeight: 2000,
      reweigh: {
        weight: 300,
      },
    },
    {
      primeActualWeight: 100,
    },
    {
      primeActualWeight: 1000,
      reweigh: {
        weight: 200,
      },
    },
    {
      primeActualWeight: 400,
      reweigh: {
        weight: 3000,
      },
    },
  ];

  it('calculates the net weight of a ppm shipment properly', () => {
    expect(calculatePPMShipmentNetWeight(ppmShipments[0])).toEqual(5000);
  });

  it('calculates the net weight of a non-ppm shipment properly', () => {
    expect(calculateNonPPMShipmentNetWeight(hhgShipments[0])).toEqual(5);
  });

  it('calculates the sum net weight of a move with varied shipment types', () => {
    const netWeightOfPPMShipments = 37000;
    const netWeightOfNonPPMShipments = 1005;

    const totalMoveWeight = [...ppmShipments, ...hhgShipments]
      .map((s) => calculateShipmentNetWeight(s))
      .reduce((accumulator, current) => accumulator + current, 0);
    expect(totalMoveWeight).toEqual(netWeightOfPPMShipments + netWeightOfNonPPMShipments);
  });
});
