import {FC, ReactNode} from 'react';

type Props = {
    title: string;
    description: string;
    children?: ReactNode;
    disableLogo?: boolean;
};

const Card: FC<Props> = ({ title, description, children, disableLogo = false}) => {
 return (
     <div className="flex flex-col gap-4 rounded-2xl p-8 bg-zinc-800 shadow-lg relative">
         {!disableLogo && <div className="w-16 h-16 bg-blue-700 rounded-full" />}
         <div className="flex flex-col gap-8">
             <div className="flex flex-col gap-2 lg:min-w-[24rem]">
                 <h1>{title}</h1>
                 <p className="text-muted-foreground">{description}</p>
             </div>
             <div className="lg:min-w-[24rem]">{children}</div>
         </div>
     </div>
 );
};

export default Card;