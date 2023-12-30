<template>
    <main>
        <div class="relative h-8">
            <div @click="search" id="icon"
                class="absolute z-10 bg-red-100 rounded-full h-full p-2 flex flex-wrap justify-center content-center hover:cursor-pointer">
                <span class="text-[white]">
                    <font-awesome-icon icon="magnifying-glass" />
                </span>
            </div>
            <div id="input" class="absolute w-full h-full">
                <input class="border-b border-b-red-100 rounded-full w-full h-full pl-10 focus-visible:outline-0"
                    v-model="searching" type="text"
                    placeholder="Buscar por remitente, destinatario o fecha">
            </div>
        </div>
    </main>
</template>

<script setup>
import { useDataStore } from '@/stores/data.js';
import { ref } from 'vue';

const emits = defineEmits([
    'search'
])

var searching = ref('')

const search = () => {
    var content = searching.value.toLowerCase()
    if ((content.search("to:") >= 0) || (content.search("from:") >= 0)) {
        var splitContent = searching.value.split(":")
        if ((splitContent[1].trim()).length >= 4) {
            emits('search', splitContent)
        }
    } else {
        //LLAMAR COMPONENTE DE ALERTA
    }
}
</script>