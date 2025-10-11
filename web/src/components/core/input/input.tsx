import {
    DetailedHTMLProps,
    FC,
    InputHTMLAttributes,
    MouseEvent,
    useId,
    useState
} from "react";
import { cn } from "../../../lib";
import { Eye, EyeOff } from "lucide-react";

type Props = DetailedHTMLProps<
    InputHTMLAttributes<HTMLInputElement>,
    HTMLInputElement
> & {
    label?: string;
    accent?: "success" | "error";
    error?: string | boolean;
};

const Input: FC<Props> = ({
    label,
    className,
    type,
    accent,
    error,
    ...props
}) => {
    const id = useId();
    const [masked, setMasked] = useState(type === "password");

    const inputType =
        type === "password" ? (masked ? "password" : "text") : type;

    const onEyeClick = (e: MouseEvent<HTMLButtonElement>) => {
        e.stopPropagation();
        e.preventDefault();
        setMasked((p) => !p);
    };

    return (
        <div className="flex flex-col gap-1.5">
            <label htmlFor={id}>{label}</label>
            <div className="relative">
                <input
                    id={id}
                    className={cn(
                        "w-full border border-border dark:bg-white/5 px-2 py-1.5 rounded outline-none focus:outline-none focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-gray-200 text-sm",
                        {
                            "border-green-500 focus-visible:ring-green-500 text-green-500 !bg-green-500/20":
                                accent === "success",
                            "border-red-500 focus-visible:ring-red-500 text-red-500":
                                error || accent === "error"
                        },
                        className
                    )}
                    type={inputType}
                    {...props}
                />
                {type === "password" && (
                    <button
                        className="p-1 absolute top-1/2 right-1 -translate-y-1/2 hover:bg-muted-foreground/20 rounded-full cursor-pointer"
                        onClick={onEyeClick}
                    >
                        {masked ? <Eye size={16} /> : <EyeOff size={16} />}
                    </button>
                )}
                {error && typeof error === "string" && (
                    <span className="text-red-500 text-sm">{error}</span>
                )}
            </div>
        </div>
    );
};

export default Input;
