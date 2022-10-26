import React from "react";
import { FiX } from "react-icons/fi";

const Modal: React.FC<
  React.PropsWithChildren<{
    title: string;
    openButon: React.ReactNode;
    buttonClass?: string;
    isOpen?: boolean;
    onOpen?: () => void;
    onClose?: () => void;
  }>
> = ({
  children,
  title,
  openButon,
  buttonClass,
  onOpen = () => {
    //no-op
  },
  onClose = () => {
    //no-op
  },
  isOpen = false,
}) => {
  return (
    <>
      {/* The button to open modal */}
      <div className={"btn " + buttonClass ?? ""} onClick={onOpen}>
        {openButon}
      </div>

      {/* Put this part before </body> tag */}

      <div className={`modal whitespace-normal ${isOpen ? "modal-open" : ""}`}>
        <div className="modal-box relative">
          <div
            className="btn btn-sm btn-circle absolute right-2 top-2 text-lg"
            onClick={onClose}
          >
            <FiX />
          </div>
          <h3 className="text-lg font-bold">{title}</h3>
          <div className="py-4">{children}</div>
        </div>
      </div>
    </>
  );
};

export default Modal;
