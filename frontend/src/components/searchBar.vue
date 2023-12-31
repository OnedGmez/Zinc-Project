<template>
    <main>
        <div class="relative h-8">
            <div @click="search" id="icon" class="absolute z-10 bg-red-100 rounded-full h-full p-2 flex flex-wrap justify-center content-center hover:cursor-pointer">
                <span class="text-[white]">
                    <font-awesome-icon icon="magnifying-glass" />
                </span>
            </div>
            <div id="input" class="absolute w-full h-full">
                <input class="border-b border-b-red-100 rounded-full w-full h-full pl-10 focus-visible:outline-0" v-model="searching" type="text" placeholder="Buscar por remitente, destinatario o asunto">
            </div>
        </div>
    </main>
    <Alert :msg="msg" v-if="alert" class="alert"/>
</template>

<script setup>
import Alert from './Alert.vue';
import { useDataStore } from '@/stores/data.js';
import { ref } from 'vue';

const emits = defineEmits([
    'search'
])

var alert = ref(false)
var searching = ref('')
var msg = ''

const search = () => {
    var content = searching.value.toLowerCase()
    if (content != '') {
        if ((content.search("to:") >= 0) || (content.search("from:") >= 0) || (content.search("subject:") >= 0)) {
            var splitContent = searching.value.split(":")
            emits('search', splitContent)
        } else {
            msg = "Antepón el criterio de búsqueda (TO:, FROM: o SUBJECT:)"
            viewAlert()
        }
    }else{
        msg = "Indica lo que deseas buscar"
        viewAlert()
    }
}

const viewAlert = () =>{
    alert.value = !alert.value
    setTimeout(() => { alert.value = !alert.value; }, 3600);
}
</script>