import { FC } from "react";

type Props = {
    value: number;
};

const Progress: FC<Props> = ({ value }) => {
    return (
        <div className="w-full h-2 rounded-full overflow-hidden bg-black/20">
            <div style={{ width: `${value}%` }} className="bg-blue-500 h-full" />
        </div>
    );
};

export default Progress;
