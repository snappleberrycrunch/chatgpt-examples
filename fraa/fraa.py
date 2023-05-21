def calculate_FRAA(putouts, assists, errors, league_putouts, league_assists, league_errors):
    """
    Calculate Fielding Runs Above Average (FRAA) for a first baseman.
    
    FRAA = (putouts - league_putouts) + (assists - league_assists) - (errors - league_errors)
    
    Args:
    putouts (int): Number of putouts by the first baseman.
    assists (int): Number of assists by the first baseman.
    errors (int): Number of errors committed by the first baseman.
    league_putouts (int): League average number of putouts for first basemen.
    league_assists (int): League average number of assists for first basemen.
    league_errors (int): League average number of errors for first basemen.
    
    Returns:
    int: Fielding Runs Above Average (FRAA).
    """
    return (putouts - league_putouts) + (assists - league_assists) - (errors - league_errors)

