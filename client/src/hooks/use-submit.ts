import {useState} from "react";

type SubmitOptions = {
    url: string;
}

const POST = async (url: string, body: string) => {
    const result = await fetch(url, { method: 'POST', body } );
    if (!result.ok || result.status > 299) {
        const errorBody = await result.json();
        throw new Error(errorBody.message ?? "Unexpected error has occurred. Please try again later.");
    }

    try {
        const data = await result.json();
        return data;
        // eslint-disable-next-line @typescript-eslint/no-unused-vars
    } catch(e: unknown) {
        return {}
    }
}

export const useSubmit = <T = unknown, D = unknown>(options: SubmitOptions) => {
    const [state, setState] = useState<'idle' | 'pending' | 'success' | 'error'>('idle');
    const [data, setData] = useState<T | undefined>();
    const [error, setError] = useState<{ message: string } | undefined>();

    console.log(state, data, error)

    const { url } = options;

    const submit =  async (data: D) => {
        setState('pending');
        try {
            const result = await POST(url, JSON.stringify(data) )
            setState('success');
            setData(result as T);
            setError(undefined);
        } catch (e: unknown) {
            const message = (e as Error)?.message ?? "Unexpected error has occurred. Please try again later."
            setError({ message });
            setState('error');
        }
    }

    return { state, data, error, submit}
}