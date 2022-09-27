import React, { createContext } from 'react';

import { addToastHandle } from '../toast/Toasts';

const ToastContext = createContext<React.MutableRefObject<addToastHandle | undefined> | null>(null);

// const ToastContext = createContext<addToastHandle>({
//   addToast: (_: string, __: toastCategory) => {
//     return;
//   },
// });
export default ToastContext;
