<template>
    <main>
        <div class="relative h-8">
            <div id="icon" class="absolute z-10 bg-red-100 rounded-full h-full p-2 flex flex-wrap justify-center content-center">
                <span class="text-[white]">
                    <font-awesome-icon icon="magnifying-glass"/>
                </span>
            </div>
            <div id="input" class="absolute w-full h-full">
                <input class="border-b border-b-red-100 rounded-full w-full h-full pl-10 focus-visible:outline-0" type="text" placeholder="Buscar por remitente, destinatario o fecha">
            </div>
        </div>
    </main>
</template>

<script setup>
</script>