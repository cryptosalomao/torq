import { useState } from 'react';

import { useGetOnChainTxQuery } from 'apiSlice';
import {
  ColumnTriple20Regular as ColumnsIcon,
  Filter20Regular as FilterIcon,
  Options20Regular as OptionsIcon,
  ArrowSortDownLines20Regular as SortIcon,
} from '@fluentui/react-icons';
import { Link } from 'react-router-dom';
import { useAppDispatch, useAppSelector } from 'store/hooks';

import useLocalStorage from 'features/helpers/useLocalStorage';
import Sidebar from 'features/sidebar/Sidebar';
import FilterSection from 'features/sidebar/sections/filter/FilterSection';
import { Clause, FilterInterface, deserialiseQuery } from 'features/sidebar/sections/filter/filter';
import SortSection, { OrderBy } from 'features/sidebar/sections/sort/SortSection';
import Table, { ColumnMetaData } from 'features/table/Table';
import Pagination from 'features/table/pagination/Pagination';
import TablePageTemplate, {
  TableControlSection,
  TableControlsButton,
  TableControlsButtonGroup,
} from 'features/templates/tablePageTemplate/TablePageTemplate';

import {
  selectActiveColumns,
  selectAllColumns,
  selectOnChainFilters,
  updateColumns,
  updateOnChainFilters,
} from './onChainSlice';
import TransactTabs from './../TransactTabs';

import { FilterCategoryType } from '../../sidebar/sections/filter/filter';
import ColumnsSection from '../../sidebar/sections/columns/ColumnsSection';
import { SectionContainer } from '../../section/SectionContainer';
import clone from '../../../clone';

type sections = {
  filter: boolean;
  sort: boolean;
  columns: boolean;
};

const statusTypes: any = {
  OPEN: 'Open',
  SETTLED: 'Settled',
  EXPIRED: 'Expired',
};

function OnChainPage() {
  const [limit, setLimit] = useLocalStorage('onchainLimit', 100);
  const [offset, setOffset] = useState(0);
  const [orderBy, setOrderBy] = useLocalStorage('onchainOrderBy', [
    {
      key: 'date',
      direction: 'desc',
    },
  ] as Array<OrderBy>);

  const activeColumns = useAppSelector(selectActiveColumns) || [];
  const allColumns = useAppSelector(selectAllColumns);

  const dispatch = useAppDispatch();
  const filters = useAppSelector(selectOnChainFilters);

  const onchainResponse = useGetOnChainTxQuery({
    limit: limit,
    offset: offset,
    order: orderBy,
    filter: filters && deserialiseQuery(filters).length >= 1 ? filters : undefined,
  });

  // Logic for toggling the sidebar
  const [sidebarExpanded, setSidebarExpanded] = useState(false);
  let data: any = [];

  if (onchainResponse?.data?.data) {
    data = onchainResponse?.data?.data.map((invoice: any) => {
      const invoice_state = statusTypes[invoice.invoice_state];

      return {
        ...invoice,
        invoice_state,
      };
    });
  }

  // General logic for toggling the sidebar sections
  const initialSectionState: sections = {
    filter: false,
    sort: false,
    columns: false,
  };

  const [activeSidebarSections, setActiveSidebarSections] = useState(initialSectionState);

  const sidebarSectionHandler = (section: keyof sections) => {
    return () => {
      setActiveSidebarSections({
        ...activeSidebarSections,
        [section]: !activeSidebarSections[section],
      });
    };
  };

  const closeSidebarHandler = () => {
    return () => {
      setSidebarExpanded(false);
    };
  };

  const tableControls = (
    <TableControlSection>
      <TransactTabs />
      <TableControlsButtonGroup>
        <TableControlsButton onClickHandler={() => setSidebarExpanded(!sidebarExpanded)} icon={OptionsIcon} />
      </TableControlsButtonGroup>
    </TableControlSection>
  );

  const defaultFilter: FilterInterface = {
    funcName: 'gte',
    category: 'number' as FilterCategoryType,
    parameter: 0,
    key: 'amount',
  };

  const filterColumns = clone(allColumns).filter(({ key }) =>
    [
      'date',
      'dest_addresses',
      'dest_addresses_count',
      'amount',
      'total_fees',
      'label',
      'lnd_tx_type_label',
      'lnd_short_chan_id',
    ].includes(key)
  );

  const handleFilterUpdate = (updated: Clause) => {
    dispatch(updateOnChainFilters({ filters: updated.toJSON() }));
  };

  const sortableColumns = allColumns.filter((column: ColumnMetaData) =>
    [
      'date',
      'dest_addresses',
      'dest_addresses_count',
      'amount',
      'total_fees',
      'label',
      'lnd_tx_type_label',
      'lnd_short_chan_id',
    ].includes(column.key)
  );

  const handleSortUpdate = (updated: Array<OrderBy>) => {
    setOrderBy(updated);
  };

  const updateColumnsHandler = (columns: Array<any>) => {
    dispatch(updateColumns({ columns: columns }));
  };

  const sidebar = (
    <Sidebar title={'Options'} closeSidebarHandler={closeSidebarHandler()}>
      <SectionContainer
        title={'Columns'}
        icon={ColumnsIcon}
        expanded={activeSidebarSections.columns}
        handleToggle={sidebarSectionHandler('columns')}
      >
        <ColumnsSection columns={allColumns} activeColumns={activeColumns} handleUpdateColumn={updateColumnsHandler} />
      </SectionContainer>
      <SectionContainer
        title={'Filter'}
        icon={FilterIcon}
        expanded={activeSidebarSections.filter}
        handleToggle={sidebarSectionHandler('filter')}
      >
        <FilterSection
          columnsMeta={filterColumns}
          filters={filters}
          filterUpdateHandler={handleFilterUpdate}
          defaultFilter={defaultFilter}
        />
      </SectionContainer>
      <SectionContainer
        title={'Sort'}
        icon={SortIcon}
        expanded={activeSidebarSections.sort}
        handleToggle={sidebarSectionHandler('sort')}
      >
        <SortSection columns={sortableColumns} orderBy={orderBy} updateHandler={handleSortUpdate} />
      </SectionContainer>
    </Sidebar>
  );

  const breadcrumbs = [
    <span key="b1">Transactions</span>,
    <Link key="b2" to={'/transactions/on-chain'}>
      On-Chain Tx
    </Link>,
  ];
  const pagination = (
    <Pagination
      limit={limit}
      offset={offset}
      total={onchainResponse?.data?.pagination?.total}
      perPageHandler={setLimit}
      offsetHandler={setOffset}
    />
  );
  return (
    <TablePageTemplate
      title={'OnChain'}
      breadcrumbs={breadcrumbs}
      sidebarExpanded={sidebarExpanded}
      sidebar={sidebar}
      tableControls={tableControls}
      pagination={pagination}
    >
      <Table
        data={data}
        activeColumns={activeColumns || []}
        isLoading={onchainResponse.isLoading || onchainResponse.isFetching || onchainResponse.isUninitialized}
      />
    </TablePageTemplate>
  );
}

export default OnChainPage;
