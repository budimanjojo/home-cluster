/*
  This is a template config file
  \{\{ .var \}\} will be replaced by `external-secrets`
*/
import { execSync } from 'child_process';

const qbittorrentURL = "http://qbittorrent.default.svc.cluster.local:8080";
const prowlarrURL = "http://prowlarr.default.svc.cluster.local:9696";
const sonarrURL = "http://sonarr.default.svc.cluster.local:8989";
const radarrURL = "http://radarr.default.svc.cluster.local:7878";

function fetchIndexers(baseUrl, apiKey, tag) {
  const buffer = execSync(`curl -fsSL "$${baseUrl}/api/v1/tag/detail?apiKey=$${apiKey}"`);
  const response = JSON.parse(buffer.toString("utf8"));
  const indexerIds = response.filter(i => i.label === tag)[0]?.indexerIds ?? [];
  const indexers = indexerIds.map(i => `$${baseUrl}/$${i}/api?apikey=$${apiKey}`);
  console.log(`Loaded $${indexers.length} indexers from Prowlarr`);
  return indexers;
}

export default {
  port: Number(process.env.CROSS_SEED_PORT),
  torznab: fetchIndexers(prowlarrURL, "{{ .PROWLARR_API_KEY }}", "cross-seed"),
  torrentClients: [`qbittorrent:$${qbittorrentURL}`],
  sonarr: [`$${sonarrURL}/?apikey={{ .SONARR_API_KEY }}`],
  radarr: [`$${radarrURL}/?apikey={{ .RADARR_API_KEY }}`],
  useClientTorrents: true,
  outputDir: null,
  linkCategory: "cross-seed",
  linkDirs: ["/downloads/cross-seed"],
  linkType: "hardlink",
  matchMode: "partial",
  includeNonVideos: false,
  action: "inject",
}
