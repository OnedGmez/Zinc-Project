<template>
    <table class="table-fixed">
        <thead>
            <tr>
                <td v-for="encabezado in encabezados" class="encabezado data-table">
                    {{ encabezado }}
                </td>
            </tr>
        </thead>
        <tbody>
            <tr v-for="i in 3" class="special-row">
                <td class="truncate data-table">
                    example@prueba.com
                </td>
                <td class="truncate data-table">
                    example@prueba.com, example@prueba.com, example@prueba.com, example@prueba.com
                </td>
                <td class="truncate data-table">
                    Thursday 15, 2023
                </td>
                <td class="truncate data-table">
                    Prueba de interfaz
                </td>
            </tr>
        </tbody>
    </table>
</template>

<script setup>
import { ref } from 'vue';

var encabezados = ["TO", "FROM", "DATE", "SUBJECT"]
</script>