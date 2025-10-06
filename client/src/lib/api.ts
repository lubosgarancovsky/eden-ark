// const postLoginCredentials = async (email: string, password: string, returnTo: string) => {
//     try {
//         const result = await fetch(`/oauth2/login?returnTo=${returnTo}`, {
//             method: "POST",
//             body: JSON.stringify({email, password}),
//         })
//
//         return result.ok;
//     } catch(e) {
//         console.error(e)
//         return false;
//     }
// }
