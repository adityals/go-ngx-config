import type { ChangeEvent } from 'react';

interface Props {
  handleChangeOpts: (name: string, e: ChangeEvent<HTMLInputElement>) => void;
}

const ParseOptions = (props: Props) => {
  return (
    <form>
      <div className="flex items-center mb-4">
        <input
          id="skip-check-ctx"
          type="checkbox"
          value=""
          onChange={(e: ChangeEvent<HTMLInputElement>) => {
            props.handleChangeOpts('skip-check-ctx', e);
          }}
          className="w-4 h-4 text-blue-600 bg-gray-100 rounded border-gray-300 focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600"
        />
        <label htmlFor="skip-check-ctx" className="ml-2 text-sm font-medium text-gray-900 dark:text-gray-300">
          Skip Check Ctx
        </label>
      </div>
    </form>
  );
};

export default ParseOptions;
