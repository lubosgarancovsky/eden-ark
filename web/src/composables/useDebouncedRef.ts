import { type Ref, type ComputedRef, ref, watch } from 'vue'

export function useDebouncedRef<T>(source: Ref<T> | ComputedRef<T>, delay = 300) {
    const debounced = ref(source.value) as Ref<T>
    let timeout: number

    watch(source, (val) => {
        clearTimeout(timeout)
        timeout = window.setTimeout(() => {
            debounced.value = val
        }, delay)
    })

    return debounced
}