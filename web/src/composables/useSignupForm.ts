import {useForm, useField} from 'vee-validate'
import {toTypedSchema} from "@vee-validate/zod";
import {useUsernameIsAvailable} from "@/composables/useUsernameIsAvailable.ts";
import {computed} from "vue";
import { z } from 'zod'

const signupSchema = toTypedSchema(
    z.object({
        firstName: z.string().min(1, "First name is required"),
        lastName: z.string().min(1, 'Last name is required'),
        username: z.string().min(1, "Username is required"),
        email: z.string().min(1, "E-mail is required"),
        password: z.string().min(8, 'Password must be at least 8 characters is required'),
        passwordRepeat: z.string().min(1, 'Please repeat your password')
    }).refine((data) => data.password === data.passwordRepeat, {
        message: "Passwords must match",
        path: ["passwordRepeat"],
    })
)


export function useSignupForm() {
    const { resetForm } = useForm({
        validationSchema: signupSchema,
        validateOnMount: false,
        initialValues: {
            firstName: '',
            lastName: '',
            username: '',
            email: '',
            password: '',
            passwordRepeat: '',
        },
    })
    const { value: firstName, errorMessage: firstNameError } = useField<string>('firstName')
    const { value: lastName, errorMessage: lastNameError } = useField<string>('lastName')
    const { value: username, errorMessage: usernameError } = useField<string>('username')
    const { value: email, errorMessage: emailError } = useField<string>('email')
    const { value: password, errorMessage: passwordError } = useField<string>('password')
    const { value: passwordRepeat, errorMessage: passwordRepeatError } = useField<string>('passwordRepeat')

    const usernameValue = computed(() => username.value)


    const isAvailable = useUsernameIsAvailable(usernameValue);

    return {
        resetForm,
        isAvailable,
        firstName: {
            value: firstName,
            error: firstNameError,
        },
        lastName: {
            value: lastName,
            error: lastNameError,
        },
        username: {
            value: username,
            error: usernameError,
        },
        email: {
            value: email,
            error: emailError,
        },
        password: {
            value: password,
            error: passwordError,
        },
        passwordRepeat: {
            value: passwordRepeat,
            error: passwordRepeatError,
        },
    }
}