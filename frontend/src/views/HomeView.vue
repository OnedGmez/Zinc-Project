<script setup>
import { ref } from 'vue';

import searchBar from '@/components/searchBar.vue';
import TableData from '@/components/TableData.vue';
import PaginationBar from '@/components/PaginationBar.vue';

var msgData = ref('')

const previewMsg = (data) => {
  msgData.value = data
  console.log(data)
}
</script>

<template>
  <main class="pt-4 h-normal">
    <section>
      <div id="search">
        <searchBar class="spacing w-10/12" />
      </div>
    </section>
    <section class="pt-4 h-normal">
      <div class="flex flex-row px-4 h-full">
        <div class="w-3/5 h-full">
          <TableData @preview="(msg) => previewMsg(msg)" class="spacing w-11/12" />
          <PaginationBar class="h-restante z-10" />
        </div>
        <div class="w-2/5">
          <p class="whitespace-break-spaces">
            {{msgData}}
          </p>
        </div>
      </div>
    </section>
  </main>
</template>
