<script setup lang="ts">
import { Form, type SubmissionHandler, useField, useForm } from "vee-validate";
import { Input } from "@/components/ui/input";
import { AdvancedInput } from "@/components";
import { ArrowLeft, RefreshCcw } from "lucide-vue-next";
import { Button } from "@/components/ui/button";
import { ref, watch } from "vue";
import { useMutation } from "@tanstack/vue-query";
import { type ApiError, requestPasswordReset } from "@/lib/api";
import { object, email as zodEmail } from 'zod'
import { toTypedSchema } from "@vee-validate/zod";

const emailSchema = toTypedSchema(object({
    email: zodEmail("This is not a valid e-mail address"),
}))

type ForgotPasswordForm = {
    email: string;
}

const { handleSubmit } = useForm<ForgotPasswordForm>({
    validationSchema: emailSchema,
    validateOnMount: false,
    initialValues: {
        email: '',
    },
})

const mutation = useMutation<void, ApiError, string>({
    mutationFn: (email) => requestPasswordReset(email),
})

const onSubmit: SubmissionHandler<any> = handleSubmit((values) => {
    mutation.mutate(values.email)
})

const { value: email, errorMessage: emailError } = useField<string>('email')

const emailWasSent = ref(false);

watch(mutation.isSuccess, (isSuccess) => emailWasSent.value = isSuccess)

const tryAgain = () => {
    emailWasSent.value = false
}
</script>

<template>
<Form class="w-full flex flex-col gap-6" @submit="onSubmit">
    <div class="flex flex-col gap-2 w-full">
        <AdvancedInput v-if="!emailWasSent" label="E-mail" :error="emailError">
            <Input v-model="email" name="email" type="email" required />
        </AdvancedInput>

        <p v-if="!emailWasSent" class="text-sm text-muted-foreground">
            We will send you an e-mail to reset your password.
        </p>
        <div v-else class="text-sm text-muted-foreground flex flex-col gap-2">
            <p>Password reset instructions were sent to:</p>
            <div class="border p-4 rounded-md text-center w-full bg-input/30 font-bold">{{ email }}</div>
        </div>


    </div>

    <div class="flex flex-col gap-4 w-full flex-1">
        <p v-if="emailWasSent" class="text-sm text-muted-foreground">
            You didn't receive the e-mail?
        </p>
        <Button v-if="!emailWasSent" type="submit" class="w-full cursor-pointer" :disabled="mutation.isPending.value">
            Send e-mail
        </Button>
        <Button v-else type="button" class="w-full cursor-pointer" @click="tryAgain">
            <RefreshCcw/>
            Try again
        </Button>
        <Button variant="secondary" class="w-full" type="button" asChild>
            <a href="/login">
                <ArrowLeft/>
                Back to login
            </a>
        </Button>
    </div>
</Form>
</template>
