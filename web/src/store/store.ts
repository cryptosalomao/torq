import { Action, configureStore, ThunkAction } from '@reduxjs/toolkit';
import { setupListeners } from '@reduxjs/toolkit/query';
import { torqApi } from 'apiSlice';

import invoicesReducer from 'features/transact/Invoices/invoicesSlice';
import onChainReducer from 'features/transact/OnChain/onChainSlice';
import paymentsReducer from 'features/transact/Payments/paymentsSlice';

import channelReducer from '../features/channel/channelSlice';
import tableReducer from '../features/forwards/forwardsSlice';
import navReducer from '../features/navigation/navSlice';
import timeIntervalReducer from '../features/timeIntervalSelect/timeIntervalSlice';

export const store = configureStore({
  reducer: {
    navigation: navReducer,
    table: tableReducer,
    payments: paymentsReducer,
    invoices: invoicesReducer,
    onChain: onChainReducer,
    timeInterval: timeIntervalReducer,
    channel: channelReducer,
    [torqApi.reducerPath]: torqApi.reducer,
  },
  middleware: (getDefaultMiddleware) => getDefaultMiddleware().concat(torqApi.middleware),
});

setupListeners(store.dispatch);

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;
export type AppThunk<ReturnType = void> = ThunkAction<ReturnType, RootState, unknown, Action<string>>;
