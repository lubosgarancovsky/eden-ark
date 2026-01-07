import { ref, computed, onMounted, onUnmounted } from 'vue'

export const useCountdown = (expiresAt: string) => {
    const now = ref(new Date())
    const expires = new Date(expiresAt)

    const totalDuration = expires.getTime() - now.value.getTime() // ms

    const remaining = ref(Math.max(expires.getTime() - Date.now(), 0)) // ms
    const percent = ref(0)

    let timer: number

    const update = () => {
        const current = Date.now()
        const diff = Math.max(expires.getTime() - current, 0)

        remaining.value = diff
        percent.value = Math.max((diff / totalDuration) * 100, 0)

        if (diff <= 0) {
            clearInterval(timer)
        }
    }

    onMounted(() => {
        timer = window.setInterval(update, 1000) // update every second
        update() // initial call
    })

    onUnmounted(() => {
        clearInterval(timer)
    })

    const formattedTime = computed(() => {
        const totalSec = Math.floor(remaining.value / 1000)
        const min = Math.floor(totalSec / 60)
        const sec = totalSec % 60
        return `${min}:${sec.toString().padStart(2, '0')}`
    })

    return {
        remaining,
        percent,
        formattedTime,
    }
}
