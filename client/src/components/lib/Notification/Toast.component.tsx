import { FastAuthIcon } from "../Icons/FastAuthIcon.component";
import { LoaderIcon } from "../Icons/LoaderIcon.component";

interface ToastComponentProps {
  onClick: () => void;
  isLoading: boolean;
}

export const ToastComponent = ({ onClick, isLoading }: ToastComponentProps) => {
  return (
    <div
      className="max-w-xs bg-white border border-gray-200 rounded-xl shadow-lg dark:bg-neutral-800 dark:border-neutral-700"
      role="alert"
      aria-labelledby="hs-toast-stack-toggle-label"
    >
      <div className="flex p-4">
        <div className="shrink-0">
          {isLoading ? <LoaderIcon /> : <FastAuthIcon />}
        </div>
        <div className="ms-4">
          <h3
            id="hs-toast-stack-toggle-label"
            className="text-gray-800 font-semibold dark:text-white"
          >
            FastAuth service
          </h3>
          <div className="mt-1 text-sm text-gray-600 dark:text-neutral-400">
            Create your account with one click and gain access
          </div>
          <div className="mt-4">
            <div className="flex gap-x-3">
              <button
                type="button"
                disabled={isLoading}
                onClick={onClick}
                className="text-blue-600 decoration-2 hover:underline font-medium text-sm focus:outline-none focus:underline dark:text-blue-500"
              >
                Go!
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};