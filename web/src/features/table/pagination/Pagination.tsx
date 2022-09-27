// import Reactfrom "react";
import styles from './pagination.module.scss';
// import classNames from "classnames";
import classNames from 'classnames';
import React, { useMemo } from 'react';
import { ChevronLeft20Filled as LeftIcon, ChevronRight20Filled as RightIcon } from '@fluentui/react-icons';
import PaginationSelect from './PaginationSelect';

type PaginationProps = {
  limit: number;
  offset: number;
  total: number;
  perPageHandler: (limit: number) => void;
  offsetHandler: (offset: number) => void;
};

const limitOptions = [
  { value: 10, label: 10 },
  { value: 20, label: 20 },
  { value: 50, label: 50 },
  { value: 100, label: 100 },
  { value: 200, label: 200 },
  { value: 500, label: 500 },
  { value: 1000, label: 1000 },
  { value: 2000, label: 2000 },
];

function renderPages(
  limit: number,
  offset: number,
  pages: number,
  currentPage: number,
  pageSelectOptions: Array<{ value: number; label: number | string }>,
  offsetHandler: (offset: number) => void
) {
  return (
    <div className={styles.paginationButtons}>
      <button
        className={classNames(styles.pageButton, { [styles.disabled]: !(offset >= limit) })}
        onClick={() => {
          if (offset >= limit) {
            offsetHandler(offset - limit);
          }
        }}
      >
        <LeftIcon />
      </button>
      <PaginationSelect
        options={pageSelectOptions}
        menuPlacement={'top'}
        className={styles.offsetSelector}
        value={{ value: offset / limit, label: `Page ${offset / limit + 1} of ${pages}` }}
        onChange={(item: any) => {
          offsetHandler(item.value * limit);
        }}
      />

      <button
        className={classNames(styles.pageButton, { [styles.disabled]: !(pages > currentPage + 1) })}
        onClick={() => {
          if (pages > currentPage + 1) {
            offsetHandler(offset + limit);
          }
        }}
      >
        <RightIcon />
      </button>
    </div>
  );
}

function Pagination(props: PaginationProps) {
  const currentPage = Math.floor(props.offset / props.limit);

  const [pages, pageSelectOptions] = useMemo(() => {
    const pages = Math.ceil(props.total / props.limit);
    const pageSelectOptions = Array.from({ length: pages }, (_, i) => {
      return { value: i, label: `${i + 1} of ${pages}` };
    });

    return [pages, pageSelectOptions];
  }, [props.total, props.limit]);

  return (
    <div className={styles.paginationContainer}>
      <div className={styles.perPageSelectWrapper}>
        <span className={styles.perPageText}>Per page: </span>
        <PaginationSelect
          options={limitOptions}
          menuPlacement={'top'}
          className={styles.limitSelector}
          value={limitOptions.find(({ value }) => value === props.limit)}
          onChange={(item: any) => {
            props.perPageHandler(item.value);
            props.offsetHandler(0);
          }}
        />
      </div>
      <div className={styles.paginationButtons}>
        {renderPages(props.limit, props.offset, pages, currentPage, pageSelectOptions, props.offsetHandler)}
      </div>
    </div>
  );
}

export default React.memo(Pagination);
