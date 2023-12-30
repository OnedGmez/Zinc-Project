import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useDataStore = defineStore('data', () => {
  var accumulatedData = ref([])
  var pages = ref(1)
  var limit = 16
  var actualP = ref(1)

  const regexpS = new RegExp(/\[.SALTO.\]+/g)
  const regexpSa = new RegExp(/\[.SANGRIA.\]+/g)

  return { accumulatedData, pages, limit, actualP, regexpS, regexpSa }
})
