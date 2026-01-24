<script setup lang="ts">
import KPIGrid from '../components/KPIGrid.vue';
import DonutChart from '../components/charts/DonutChart.vue';
import LineChart from '../components/charts/LineChart.vue';
import BillsTable from '../components/BillsTable.vue';
</script>

<template>
  <KPIGrid />
  
  <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
    <DonutChart />
    <LineChart />
  </div>

  <BillsTable />
</template>
