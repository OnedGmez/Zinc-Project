<template>
    <div class="flex items-center justify-between border-t border-red-100 bg-white px-10 py-3">
        <div class="flex flex-1 justify-between sm:hidden">
            <a href="#"
                class="relative inline-flex items-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50">Previous</a>
            <a href="#"
                class="relative ml-3 inline-flex items-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50">Next</a>
        </div>
        <div class="hidden sm:flex sm:flex-1 sm:items-center sm:justify-between">
            <div>
                <p class="text-sm text-gray-700">
                    Página
                    <span class="font-medium">{{ store.actualP }}</span>
                    de
                    <span class="font-medium"> {{ store.pages }} </span>
                </p>
            </div>
            <div class="flex">
                <!--BOTONES-->
                <div id="prev" @click="changeDownPage('m')" disabled class="mr-7 hover:cursor-pointer hover:text-red-100">
                    <span>
                        <font-awesome-icon icon="angles-left" />
                    </span>
                </div>
                <div id="next" @click="changeUpPage('p')" class="ml-7 hover:cursor-pointer hover:text-red-100">
                    <span>
                        <font-awesome-icon icon="angles-right" />
                    </span>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, watchEffect } from 'vue';
import { useDataStore } from '@/stores/data.js';

const store = useDataStore();

const changeUpPage = (op) => {
    if ((op == "p") && (store.actualP != store.pages)) {
        store.actualP += 1;
    }
}

const changeDownPage = (op) => {
    if (store.actualP > 1) {
        store.actualP -= 1;
    }
}

watchEffect(() => {
    if (store.accumulatedData != '') {
        store.actualP = 1
    }
})
</script>