export const ModalBody = ({ children }: { children: string }) => {
  return (
    <>
      <div className="relative p-6 flex-auto">
        <p className="my-4 text-blueGray-500 text-lg leading-relaxed">
          {children}
        </p>
      </div>
    </>
  );
};
