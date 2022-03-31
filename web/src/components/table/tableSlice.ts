import { createAsyncThunk, createSlice, PayloadAction } from '@reduxjs/toolkit';
import { RootState, AppThunk } from '../../store/store';
import {log} from "util";
import {addDays, format} from 'date-fns';
import {FilterInterface, FilterFunctions, applyFilters} from './filter'


export interface ColumnMetaData {
  heading: string;
  key: string;
  type?: string;
  width?: number;
  locked?: boolean;
  valueType: string;
}

export const columns: ColumnMetaData[] = [
  { heading: "Name", type: "AliasCell", key: "alias", locked: true, valueType: "string"},
  { heading: "Revenue", type: "BarCell", key: "revenue_out", valueType: "number"},
  { heading: "Successful outbound", type: "BarCell", key: "count_out", valueType: "number"},
  { heading: "Successful inbound", type: "BarCell", key: "count_in", valueType: "number"},
  { heading: "Successful total", type: "BarCell", key: "count_total", valueType: "number"},
  { heading: "Amount outbound", type: "BarCell", key: "amount_out", valueType: "number"},
  { heading: "Amount inbound", type: "BarCell", key: "amount_in", valueType: "number"},
  { heading: "Amount total", type: "BarCell", key: "amount_total", valueType: "number"},
  { heading: "Contributed (revenue inbound)", type: "BarCell", key: "revenue_in", valueType: "number"},
  { heading: "Contributed (revenue total)", type: "BarCell", key: "revenue_total", valueType: "number"},
  { heading: "Turnover outbound", type: "NumericCell", key: "turnover_out", valueType: "number"},
  { heading: "Turnover inbound", type: "NumericCell", key: "turnover_in", valueType: "number"},
  { heading: "Turnover total", type: "NumericCell", key: "turnover_total", valueType: "number"},
  { heading: "Capacity", type: "NumericCell", key: "capacity", valueType: "number"},
]

export interface TableState {
  channels: [];
  modChannels: [];
  filters: Array<FilterInterface>;
  columns: ColumnMetaData[];
  status: 'idle' | 'loading' | 'failed';
}

const initialState: TableState = {
  channels: [],
  modChannels: [],
  filters: loadTableState() || [],
  columns: columns,
  status: 'idle',
};
const init: RequestInit = {
  credentials: 'include',
  headers: {'Content-Type':'application/json'},
  mode: 'cors',
};

function fetchChannels(from: string, to: string) {
  to = format(addDays(new Date(to), 1), "yyyy-MM-dd")
  const body = fetch(`http://localhost:8080/api/channels?from=${from}&to=${to}`,init)
    .then(response => {
      return response.json()
    })
  return body
}

// The function below is called a thunk and allows us to perform async logic. It
// can be dispatched like a regular action: `dispatch(incrementAsync(10))`. This
// will call the thunk with the `dispatch` function as the first argument. Async
// code can then be executed and other actions can be dispatched. Thunks are
// typically used to make async requests.
export const fetchChannelsAsync = createAsyncThunk(
  'table/fetchChannels',
  async (data: {from: string, to: string}) => {
    const response = await fetchChannels(data.from, data.to);
    return response
  }
);


export function loadTableState() {
  try {
    const serializedState = localStorage.getItem("torq_table_filters");
    if (!serializedState) return undefined;
    return JSON.parse(serializedState);
  } catch (e) {
    return undefined;
  }
}

export async function saveTableState(state: any) {
  try {
    const serializedState = JSON.stringify(state);
    localStorage.setItem("torq_table_filters", serializedState);
  } catch (e) {
    console.log(e)
  }
}

export const tableSlice = createSlice({
  name: 'table',
  initialState,
  // The `reducers` field lets us define reducers and generate associated actions
  reducers: {
    updateFilters: (state, actions: PayloadAction<{filters: FilterInterface[]}>) => {
      state.filters = actions.payload.filters
      saveTableState(state.filters)
    },
  },
  // The `extraReducers` field lets the slice handle actions defined elsewhere,
  // including actions generated by createAsyncThunk or in other slices.
  extraReducers: (builder) => {
    builder
      .addCase(fetchChannelsAsync.pending, (state) => {
        state.status = 'loading';
      })
      .addCase(fetchChannelsAsync.fulfilled, (state, action) => {
        state.status = 'idle';
        state.channels = action.payload
      });
  },
});

export const { updateFilters } = tableSlice.actions;

export const selectChannels = (state: RootState) => {
  return applyFilters(state.table.filters, state.table.channels)
};
export const selectColumns = (state: RootState) => state.table.columns;
export const selectFilters = (state: RootState) => state.table.filters;

export default tableSlice.reducer;
