<template>
    <table class="table-fixed h-normal">
        <thead>
            <tr>
                <td v-for="encabezado in encabezados" class="encabezado data-table">
                    {{ encabezado }}
                </td>
            </tr>
        </thead>
        <tbody>
            <tr v-for="i in 15" @click="previewMsg()" class="special-row">
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

var msg = "Lorem ipsum, dolor sit amet consectetur adipisicing elit. \n\nQuisquam delectus, quidem magni, ipsum dicta obcaecati odit dolores expedita consequatur deleniti, vero ratione dolor atque ut mollitia in possimus hic asperiores quos quam cumque unde debitis repudiandae? Mollitia, tenetur voluptatibus magni accusamus quidem porro ipsum inventore aperiam, officia similique aliquid cum?"
var encabezados = ["TO", "FROM", "DATE", "SUBJECT"]

const emits = defineEmits([
    'preview'
])

const previewMsg = () => {
    emits('preview', msg)
}

</script>