<template>
  <main class="pt-4 h-normal">
    <section>
      <div id="search">
        <searchBar class="spacing w-10/12" @search="(splitContent) => chargeData(splitContent)" />
      </div>
    </section>
    <section class="pt-4 h-normal">
      <div class="flex flex-row px-2 h-full">
        <div :class="[msgData != '' ? 'w-3/5' : 'w-full']" class="h-full px-2 ease-in-out duration-500">
          <TableData @preview="(msg) => (msgData = msg)" class="spacing w-11/12" />
        </div>
        <div v-if="msgData != ''" class="w-2/5 px-2 overflow-y-scroll overflow-x-hidden">
          <p class="whitespace-break-spaces text-xs">
            {{ msgData }}
          </p>
        </div>
      </div>
    </section>
  </main>
</template>


<script setup>
import { ref } from 'vue';
import { useDataStore } from '@/stores/data.js';
import axios from 'axios'

import searchBar from '@/components/searchBar.vue';
import TableData from '@/components/TableData.vue';

const store = useDataStore();
var msgData = ref('')

const chargeData = async (data) => {
  try {
    axios.get('http://localhost:3000/emails', {
      params: {
        "criterion": data[0],
        "value": data[1],
      }
    }).then(response => {
      store.accumulatedData = response.data["hits"]
      store.pages = Math.ceil(response.data["total"]/15)
    })
  } catch (error) {
    console.log(error)
  }
}

</script>
