import {DetailedHTMLProps, FC, InputHTMLAttributes, useId} from 'react';
import {cn} from "../../../lib";

type Props = DetailedHTMLProps<InputHTMLAttributes<HTMLInputElement>, HTMLInputElement> & {
 label?: string
};

const Input: FC<Props> = ({label, className, ...props}) => {
    const id = useId()

     return (
          <div className='flex flex-col gap-1.5'>
              <label htmlFor={id}>{label}</label>
              <input id={id} className={cn("border border-neutral-700 p-4 rounded-lg outline-none focus:outline-none focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-gray-200", className)} {...props}/>
          </div>
     );
};

export default Input;