<template>
  <div>
    <div class="sp-container">
      <sp-sign-in />
      <sp-bank-balances />
      <sp-token-send />
      <!-- this line is used by starport scaffolding # 4 -->
		<sp-type-form type="bodyIndex" :fields="['age', 'sex', 'nation', 'weight', 'height', 'weightIndex', 'obesityWaistline', 'waistline', 'maxBloodPressure', 'minBloodPressure', 'goodCholesterol', 'batCholesterol', 'totalCholesterol', 'Dyslipidemia', 'pvd', 'sportActivities', 'education', 'marry', 'income', 'sourceCase', 'visionBad', 'drink', 'highBloodPressure', 'familialHighBloodPressure', 'diabetes', 'familialDiabetes', 'hepatitis', 'familialHepatitis', 'chronicFatigue', ]" module="woodpecker" />
		<sp-type-form type="medicalHistory" :fields="['medicalInstitutionName', 'diagnosisTime', 'diagnosisDepartment', 'diagnosisType', 'diagnosisDoctor', 'diagnosisResult', 'treatmentOptions', ]" module="woodpecker" />
    </div>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";

export default {
  components: { ...sp },
};
</script>
