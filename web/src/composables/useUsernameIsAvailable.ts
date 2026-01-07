import {useQuery} from "@tanstack/vue-query";
import {type ApiError, isUsernameAvailable} from "@/lib/api";
import {computed, type ComputedRef} from "vue";
import {useDebouncedRef} from "@/composables/useDebouncedRef.ts";

export const useUsernameIsAvailable = (usernameValue: ComputedRef<string>) => {
    const username = useDebouncedRef(usernameValue, 500)

    const usernameQuery = useQuery<Boolean, ApiError>({
        queryKey:  computed(() => ['username-available', username.value]),
        queryFn: () => isUsernameAvailable(username.value),
        enabled: computed(() => Boolean(username.value)),
    })

    return computed(() => ({
        isUsernameAvailable: usernameQuery.data.value ?? true,
        isLoading: usernameQuery.isFetching,
    }))
};