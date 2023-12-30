<template>
    <div class="h-normal w-full">
        <table :class="[i >= 16 ? 'h-full': 'h-auto']" class="table-fixed w-full">
            <thead>
                <tr>
                    <td v-for="encabezado in encabezados" class="encabezado data-table">
                        {{ encabezado }}
                    </td>
                </tr>
            </thead>
            <tbody>
                <tr v-if="data != ''" v-for="email in dataPages[store.actualP - 1]" @click="previewMsg(email.body)" class="special-row">
                    <td class="truncate data-table">
                        {{ email.to }}
                    </td>
                    <td class="truncate data-table">
                        {{ email.from }}
                    </td>
                    <td class="truncate data-table">
                        {{ email.date }}
                    </td>
                    <td class="truncate data-table">
                        {{ email.subject }}
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
    <div class="h-restante ">
        <PaginationBar class="z-10" />
    </div>
</template>

<script setup>
import PaginationBar from '@/components/PaginationBar.vue';

import { useDataStore } from '@/stores/data.js';
import { ref, watchEffect } from 'vue';

var msg = ''
var data = ref([])
var dataPages = ref([])
var i = 1
var encabezados = ["TO", "FROM", "DATE", "SUBJECT"]
const store = useDataStore()

const emits = defineEmits([
    'preview'
])

const previewMsg = (data) => {
    msg = data
    msg = msg.replace(store.regexpS, '\n').replace(store.regexpSa, '\t')
    emits("preview", msg)
}

watchEffect(() => {
    if (store.accumulatedData != '') {
        msg = ''
        emits("preview", msg)
        dataPages.value = []
        i = 1
        var dataTMP = []
        data.value = store.accumulatedData

        data.value.forEach(row => {
            if (i % store.limit != 0) {
                dataTMP.push(row)
            } else {
                dataPages.value.push(dataTMP)
                dataTMP = []
                dataTMP.push(row)
            }

            if (i == data.value.length && dataTMP != []){
                dataPages.value.push(dataTMP)
                dataTMP = []
            }
            i++
        });
    }
})

</script>