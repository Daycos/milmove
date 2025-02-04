import React, { useState } from 'react';
import { Formik } from 'formik';
import * as Yup from 'yup';
// import { func, number, string, bool } from 'prop-types';
import { Button, Fieldset, Label, Textarea } from '@trussworks/react-uswds';

import styles from './EditPPMNetWeight.module.scss';

import MaskedTextField from 'components/form/fields/MaskedTextField/MaskedTextField';
import { ErrorMessage } from 'components/form/ErrorMessage';
import { formatWeight } from 'utils/formatters';
import { calculateNetWeightForWeightTicket } from 'utils/shipmentWeights';
import { useCalculatedWeightRequested } from 'hooks/custom';

// Labels & constants

const CALCULATION_TYPE = {
  NET_WEIGHT: 'NET_WEIGHT',
  EXCESS_WEIGHT: 'EXCESS_WEIGHT',
  REDUCE_WEIGHT: 'REDUCE_WEIGHT',
};
const weightLabels = {
  [CALCULATION_TYPE.NET_WEIGHT]: {
    firstLabel: ' | original weight',
    secondLabel: ' | to fit within weight allowance',
  },
  [CALCULATION_TYPE.REDUCE_WEIGHT]: {
    firstLabel: ' | original weight',
    secondLabel: ' | to reduce excess weight',
  },
  [CALCULATION_TYPE.EXCESS_WEIGHT]: {
    firstLabel: 'Move weight (total)',
    secondLabel: 'Weight allowance',
    thirdLabel: 'Excess weight (total)',
  },
};

// Flexbox wrapper
const FlexContainer = ({ children, className }) => {
  return (
    <div className={className} style={{ display: 'flex' }}>
      {children}
    </div>
  );
};

// Form Error Indicator
const ErrorIndicator = ({ children, hasErrors }) => {
  return (
    <div data-testid="errorIndicator" className={hasErrors ? 'usa-form-group--error' : ''}>
      {children}
    </div>
  );
};

const WeightCalculationHint = ({ type, firstValue, secondValue, thirdValue }) => {
  const { firstLabel, secondLabel, thirdLabel } = weightLabels[type];
  return (
    <>
      <FlexContainer className={styles.minus}>
        {thirdValue && <>-</>}
        <div className={styles.calculationWrapper}>
          <div className={styles.calculations}>
            <strong className={styles.value}>{formatWeight(firstValue)}</strong>
            <span className={styles.label}> {firstLabel}</span>
          </div>
          {secondValue && (
            <div className={styles.calculations}>
              <strong className={styles.value}>{formatWeight(secondValue)}</strong>
              <span className={styles.label}> {secondLabel}</span>
            </div>
          )}
        </div>
      </FlexContainer>
      {thirdValue && (
        <>
          <hr className={styles.divider} />
          <div className={styles.calculations}>
            <strong className={styles.value}>{formatWeight(thirdValue)}</strong>
            <span className={styles.label}> {thirdLabel}</span>
          </div>
        </>
      )}
    </>
  );
};

const validationSchema = Yup.object({
  ppmNetWeight: Yup.number().min(0, 'Net weight must be 0 lbs or greater').required('Required'),
  ppmNetWeightRemarks: Yup.string().required('Required'),
});

const EditPPMNetWeightForm = ({ onSave, onCancel, initialValues }) => (
  <Formik initialValues={initialValues} validationSchema={validationSchema}>
    {({ handleChange, values, isValid, errors, touched, setTouched }) => (
      <div>
        <Fieldset className={styles.fieldset}>
          <MaskedTextField
            data-testid="weightInput"
            defaultValue="0"
            id="ppmNetWeight"
            name="ppmNetWeight"
            mask={Number}
            lazy={false}
            scale={0}
            signed={false} // no negative numbers
            thousandsSeparator=","
            suffix="lbs"
            inputClassName={styles.weightInput}
            errorClassName={styles.errors}
            labelClassName={styles.weightLabel}
          />
          <Label htmlFor="remarks">Remarks</Label>
          <ErrorMessage
            className={styles.errors}
            display={!!touched.ppmNetWeightRemarks && !!errors.ppmNetWeightRemarks}
          >
            {errors.ppmNetWeightRemarks}
          </ErrorMessage>
          <ErrorIndicator hasErrors={!!touched.ppmNetWeightRemarks && !!errors.ppmNetWeightRemarks}>
            <Textarea
              id="ppmNetWeightRemarks"
              data-testid="formRemarks"
              maxLength={500}
              placeholder=""
              onChange={handleChange}
              onBlur={() => {
                setTouched({ ppmNetWeightRemarks: true }, false);
              }}
              value={values.ppmNetWeightRemarks}
            />
          </ErrorIndicator>
          <FlexContainer className={styles.wrapper}>
            <Button onClick={onSave} disabled={!isValid}>
              Save changes
            </Button>
            <Button unstyled onClick={onCancel}>
              Cancel
            </Button>
          </FlexContainer>
        </Fieldset>
      </div>
    )}
  </Formik>
);

const EditPPMNetWeight = ({ netWeightRemarks, weightTicket, weightAllowance, shipments }) => {
  const [showEditForm, setShowEditForm] = useState(false);

  // To-do: Add mutation and onSumbit handler

  const toggleEditForm = () => {
    setShowEditForm(!showEditForm);
  };
  // Original weight is the full weight - empty weight
  const originalWeight = calculateNetWeightForWeightTicket(weightTicket);
  // moveWeightTotal = Sum of all ppm weights + sum of all non-ppm shipments
  const moveWeightTotal = useCalculatedWeightRequested(shipments);
  const excessWeight = moveWeightTotal - weightAllowance;
  const hasExcessWeight = Boolean(excessWeight > 0);
  // To-do: Once new netWeights can be saved, the toFit value should use that value here if its availble over using original weight
  const netWeight = originalWeight;
  const toFitValue = hasExcessWeight ? -Math.min(excessWeight, netWeight) : null;
  const showWarning = Boolean(hasExcessWeight && !showEditForm);
  const showReduceWeight = Boolean(-originalWeight === toFitValue);
  return (
    <div className={styles.wrapper}>
      <div>
        <h4 className={styles.mainHeader}>Edit PPM net weight</h4>
        {Boolean(showEditForm && hasExcessWeight) && (
          <WeightCalculationHint
            firstValue={moveWeightTotal}
            secondValue={weightAllowance}
            thirdValue={excessWeight}
            type={CALCULATION_TYPE.EXCESS_WEIGHT}
          />
        )}
      </div>
      <FlexContainer className={styles.netWeightContainer}>
        {showWarning && <div className={styles.warnings} data-testid="warning" />}
        <div>
          <h5 className={styles.header}>Net weight</h5>
          <WeightCalculationHint
            firstValue={originalWeight}
            secondValue={toFitValue}
            type={showReduceWeight ? CALCULATION_TYPE.REDUCE_WEIGHT : CALCULATION_TYPE.NET_WEIGHT}
          />
          {!showEditForm ? (
            <div className={styles.wrapper}>
              {formatWeight(netWeight)}
              {netWeightRemarks && (
                <>
                  <h5 className={styles.remarksHeader}>Remarks</h5>
                  <p className={styles.remarks}>{netWeightRemarks}</p>
                </>
              )}
              <Button onClick={toggleEditForm} className={styles.editButton}>
                Edit
              </Button>
            </div>
          ) : (
            <EditPPMNetWeightForm
              initialValues={{ ppmNetWeight: String(netWeight), netWeightRemarks }}
              onCancel={toggleEditForm}
            />
          )}
        </div>
      </FlexContainer>
    </div>
  );
};

export default EditPPMNetWeight;
