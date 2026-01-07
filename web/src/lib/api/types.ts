export type ApiError = {
    code: string;
    message: string;
    correlationId: string;
    serviceId: string;
    timestamp: string;
};

export type SignUpRequest = {
    firstName: string;
    lastName: string;
    username: string;
    email: string;
    password: string;
};