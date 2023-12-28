<script setup>
import searchBar from '@/components/searchBar.vue';
import TableData from '@/components/TableData.vue'
</script>

<template>
  <main>
    <div id="search">
      <searchBar class="mt-7 spacing w-10/12" />
    </div>
    <section class="mt-8">
      <div class="flex flex-row px-4">
        <div class="w-3/5">
          <TableData class="spacing w-11/12"/>
        </div>
        <div class="w-2/5">
          
        </div>
      </div>
    </section>
  </main>
</template>
