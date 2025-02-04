/**
 * office test fixture for playwright
 * Using javascript classes following the examples in
 * <https://playwright.dev/docs/test-fixtures>
 */
// @ts-check
import * as base from '@playwright/test';

import { BaseTestPage } from './baseTest';

/**
 * devlocal auth user types
 */
export const TOOOfficeUserType = 'TOO office';
export const TIOOfficeUserType = 'TIO office';
export const QAECSROfficeUserType = 'QAE/CSR office';
export const ServicesCounselorOfficeUserType = 'Services Counselor office';
export const PrimeSimulatorUserType = 'Prime Simulator office';

/**
 * office test fixture for playwright
 * See https://playwright.dev/docs/test-fixtures
 * @extends BaseTestPage
 */
export class OfficePage extends BaseTestPage {
  /**
   * Wait for the page to finish loading.
   */
  async waitForLoading() {
    // make sure we have signed in
    await base.expect(this.page.locator('button:has-text("Sign out")')).toHaveCount(1, { timeout: 10000 });
    await base.expect(this.page.locator('h2[data-name="loading-placeholder"]')).toHaveCount(0);
  }

  /**
   * Wait for office user Move details page
   */
  async waitForMoveDetailsPage() {
    await this.waitForLoading();
    await base.expect(this.page.getByRole('heading', { name: 'Move details' })).toBeVisible();
  }

  /**
   * Goto URL and wait for the page to finish loading.
   *
   * @param {string} url to visit
   */
  async gotoAndWaitForLoading(url) {
    await this.page.goto(url);
    await this.waitForLoading();
  }

  /**
   * Sign in as existing office user with devlocal
   *
   * @param {string} email
   */
  async signInAsExistingOfficeUser(email) {
    await this.page.goto('/devlocal-auth/login');
    await this.page.locator('input[name=email]').fill(email);
    await this.page.locator('p', { hasText: 'User Email' }).locator('button').click();
  }

  /**
   * Use devlocal auth to sign in as new Service Counselor
   */
  async signInAsNewServicesCounselorUser() {
    await Promise.all([
      // It is important to call waitForNavigation before click to set up waiting.
      this.page.waitForNavigation(),
      await this.signInAsNewUser(ServicesCounselorOfficeUserType),
    ]);
    await this.waitForLoading();
  }

  /**
   * Use devlocal auth to sign in as new TIO
   */
  async signInAsNewTIOUser() {
    await Promise.all([
      // It is important to call waitForNavigation before click to set up waiting.
      this.page.waitForNavigation(),
      await this.signInAsNewUser(TIOOfficeUserType),
    ]);
    await this.waitForLoading();
  }

  /**
   * Use devlocal auth to sign in as new TOO
   */
  async signInAsNewTOOUser() {
    await Promise.all([
      // It is important to call waitForNavigation before click to set up waiting.
      this.page.waitForNavigation(),
      await this.signInAsNewUser(TOOOfficeUserType),
    ]);
    await this.waitForLoading();
  }

  /**
   * Use devlocal auth to sign in as office user with both TOO and TIO roles
   */
  async signInAsNewTIOAndTOOUser() {
    const user = await this.testHarness.buildOfficeUserWithTOOAndTIO();
    await Promise.all([
      // It is important to call waitForNavigation before click to set up waiting.
      this.page.waitForNavigation(),
      this.signInAsExistingOfficeUser(user.login_gov_email),
    ]);
    await this.waitForLoading();
  }

  /**
   * Use devlocal auth to sign in as office user with prime simulator role
   */
  async signInAsNewPrimeSimulatorUser() {
    await Promise.all([
      // It is important to call waitForNavigation before click to set up waiting.
      this.page.waitForNavigation(),
      await this.signInAsNewUser(PrimeSimulatorUserType),
    ]);
    await this.waitForLoading();
  }

  /**
   * Use devlocal auth to sign in as office user with qaecsr role
   */
  async signInAsNewQAECSRUser() {
    await Promise.all([
      // It is important to call waitForNavigation before click to set up waiting.
      this.page.waitForNavigation(),
      await this.signInAsNewUser(QAECSROfficeUserType),
    ]);
    await this.waitForLoading();
  }

  /**
   * search for and navigate to move
   * @param {string} moveLocator
   */
  async qaeCsrSearchForAndNavigateToMove(moveLocator) {
    await this.page.locator('input[name="searchText"]').type(moveLocator);
    await this.page.locator('input[name="searchText"]').blur();

    await this.page.getByRole('button', { name: 'Search' }).click();
    await this.waitForLoading();

    await base.expect(this.page.locator('tbody >> tr')).toHaveCount(1);
    await base.expect(this.page.locator('tbody >> tr').first()).toContainText(moveLocator);

    // click result to navigate to move details page
    await this.page.locator('tbody > tr').first().click();
    await this.waitForLoading();

    base.expect(this.page.url()).toContain(`/moves/${moveLocator}/details`);
  }

  /**
   * TIO search for and navigate to move
   * @param {string} moveLocator
   */
  async tioNavigateToMove(moveLocator) {
    await this.page.locator('#locator').type(moveLocator);
    await this.page.locator('th[data-testid="locator"]').first().click();
    await this.page.locator('[data-testid="locator-0"]').click();
    await this.waitForLoading();
  }

  /**
   * TOO search for and navigate to move
   * @param {string} moveLocator
   */
  async tooNavigateToMove(moveLocator) {
    await this.page.locator('input[name="locator"]').type(moveLocator);
    await this.page.locator('input[name="locator"]').blur();

    // click result to navigate to move details page
    await this.page.locator('tbody > tr').first().click();
    await this.waitForLoading();
    base.expect(this.page.url()).toContain(`/moves/${moveLocator}/details`);
  }
}

/**
 * @typedef {object} OfficePageTestArgs - office page test args
 * @property {OfficePage} officePage    - office page
 */

/** @type {base.Fixtures<OfficePageTestArgs, {}, base.PlaywrightTestArgs, base.PlaywrightWorkerArgs>} */
const officeFixtures = {
  officePage: async ({ page, request }, use) => {
    const officePage = new OfficePage(page, request);
    await use(officePage);
  },
};

export const test = base.test.extend(officeFixtures);

export const { expect } = base;

export default test;
