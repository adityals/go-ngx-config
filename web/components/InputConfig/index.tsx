import type { ChangeEvent } from 'react';

interface Props {
  name: string;
  onChange: (e: ChangeEvent<HTMLTextAreaElement>) => void;
}

const InputConfig = (props: Props) => {
  return (
    <label className="h-screen relative block ">
      <textarea
        className="h-full w-full placeholder:text-slate-100 block bg-slate-500 text-slate-100 py-2 pl-3 pr-3 sm:text-sm"
        placeholder="Type your config..."
        name={props.name}
        onChange={props.onChange}
      />
    </label>
  );
};

export default InputConfig;
