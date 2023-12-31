<template>
    <div class="flex">
        <span class="block mr-3"><font-awesome-icon icon="circle-info" /></span>
        <span class="block font-semibold"> {{ msg }} </span>
    </div>
</template>

<script setup>
const props = defineProps([
    'msg'
])
</script>