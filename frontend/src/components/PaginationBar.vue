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
                    Showing
                    <span class="font-medium">{{ pag }}</span>
                    to
                    <span class="font-medium">10</span>
                    of
                    <span class="font-medium">97</span>
                    results
                </p>
            </div>
            <div class="flex">
                <!--BOTONES-->
                <div id="prev" @click="cambioPag('m')" disabled class="mr-7 hover:cursor-pointer hover:text-red-100">
                    <span>
                        <font-awesome-icon icon="angles-left" />
                    </span>
                </div>
                <div id="next" @click="cambioPag('s')" class="ml-7 hover:cursor-pointer hover:text-red-100">
                    <span>
                        <font-awesome-icon icon="angles-right" />
                    </span>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue';
var pag = ref(1)

const cambioPag = (op) => {
    if (op == "s") {
        pag.value += 1;
    } else {
        if (pag.value > 1) {
            pag.value -= 1;
        }
    }
}
</script>