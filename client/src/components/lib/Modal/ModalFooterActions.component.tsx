import { Button } from "../Button/Button.component";
import { LoaderIcon } from "../Icons/LoaderIcon.component";

interface ModalFooterProps {
  isLoading: boolean;
  secondaryButtonCopy: string;
  primaryButtonCopy: string;
  handlePrimaryAction: () => void;
  handleSecondaryAction: () => void;
  className?: string;
}

export const ModalFooterActions = ({
  className,
  isLoading,
  primaryButtonCopy,
  secondaryButtonCopy,
  handlePrimaryAction,
  handleSecondaryAction,
}: ModalFooterProps) => {
  return (
    <>
      <div className={`flex items-center justify-end rounded-b ${className}`}>
        <Button
          variant="secondary"
          onClick={handleSecondaryAction}
          className="mr-4"
        >
          {secondaryButtonCopy}
        </Button>
        <Button
          type="submit"
          variant="primary"
          disabled={isLoading}
          onClick={handlePrimaryAction}
        >
          {isLoading && <LoaderIcon />} {primaryButtonCopy}
        </Button>
      </div>
    </>
  );
};
